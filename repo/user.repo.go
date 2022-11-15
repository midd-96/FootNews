package repo

import (
	"NewsAppApi/model"
	"NewsAppApi/utils"
	"database/sql"
	"errors"
	"log"
)

type UserRepository interface {
	ReportPost(report model.Report) error
	MarkVote(newVote model.Vote) (error, string)
	AllPosts(pagenation utils.Filter, sortby string) ([]model.PostResponse, utils.Metadata, error)
	FindUser(email string) (model.UserResponse, error)
	InsertUser(user model.User) (int, error)
	AddPost(post model.Post) (int, error)
	AddComment(comment model.Comment) error
	StoreVerificationDetails(email string, code int) error
	VerifyAccount(email string, code int) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}
func (c *userRepo) ReportPost(report model.Report) error {
	counter := 0
	query := `INSERT INTO reports(
		created_at,
		user_id,
		post_id)
		VALUES($1, $2, $3);`
	query1 := `SELECT COUNT(*) as count
				from reports 
				WHERE user_id = $1;`

	err := c.db.QueryRow(query, report.CreatedAt, report.UserID, report.PostID)
	c.db.QueryRow(query1, report.UserID).Scan(&counter)
	if counter >= 3 {
		query1 = `UPDATE users
					SET activated = false
					WHERE id= $1;`

		_ = c.db.QueryRow(query1, report.UserID)
	}

	if err == nil {
		return errors.New("Error while marking vote")
	}
	return nil
}

func (c *userRepo) MarkVote(newVote model.Vote) (error, string) {
	query := `SELECT * FROM votes 
				WHERE user_id = $1 
				AND post_id = $2;`
	res := c.db.QueryRow(query, newVote.UserID, newVote.PostID).Scan(&newVote.CreatedAt)
	if res == sql.ErrNoRows {
		log.Println("Newly voting")
		query = `INSERT INTO votes(
					created_at,
					user_id,
					post_id)
					VALUES($1, $2, $3);`
		err := c.db.QueryRow(query, newVote.CreatedAt, newVote.UserID, newVote.PostID)
		log.Println("error :", err)
		if err == nil {
			return errors.New("Error while marking vote"), "Failed"
		}
	} else {
		log.Println("already voted voting")
		query = `DELETE FROM votes 
					WHERE user_id = $1 
					AND post_id = $2;`
		err := c.db.QueryRow(query, newVote.UserID, newVote.PostID).Scan(&newVote.CreatedAt)
		if err == nil {
			return errors.New("While unmarking vote"), "Failed"
		}
		return nil, "Unvoted"
	}
	return nil, "Voted"
}

func (c *userRepo) AddComment(comment model.Comment) error {

	query := `INSERT INTO comments(
				created_at,
				body,
				post_id,
				user_id)
				VALUES 
				($1, $2, $3, $4);`

	err := c.db.QueryRow(query,
		comment.CreatedAt,
		comment.Body,
		comment.PostID,
		comment.UserID)
	log.Println("error : ", err)
	if err == nil {
		return errors.New("Failed to insert comment")
	}
	return nil
}

func (c *userRepo) AddPost(post model.Post) (int, error) {
	var id int

	query := `INSERT INTO posts(
			created_at,
			title,
			url,
			user_id
			)
			VALUES
			($1, $2, $3, $4)
			RETURNING id;`

	err := c.db.QueryRow(query,
		post.CreatedAt,
		post.Title,
		post.Url,
		post.UserID).Scan(
		&id,
	)
	return id, err

}

func (c *userRepo) AllPosts(pagenation utils.Filter, sortby string) ([]model.PostResponse, utils.Metadata, error) {

	var posts []model.PostResponse
	log.Println("Pagenation :", pagenation)
	query := `
	SELECT COUNT(*) OVER() AS total_records, pq.*, u.name as uname FROM (
	    SELECT p.id, p.title, p.url, p.created_at, p.user_id as uid, COUNT(c.post_id) as comment_count, count(v.post_id) as votes, p.approved
		FROM posts p 
		LEFT JOIN comments c ON p.id = c.post_id 
	    LEFT JOIN votes v ON p.id = v.post_id
		where p.approved = true
		GROUP BY p.id

		) AS pq
	LEFT JOIN users u ON u.id = uid`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	var totalRecords int

	defer rows.Close()

	for rows.Next() {
		var post model.PostResponse

		err = rows.Scan(
			&post.TotalRecords,
			&post.ID,
			&post.Title,
			&post.Url,
			&post.CreatedAt,
			&post.UserID,
			&post.CommentCount,
			&post.Votes,
			&post.Approved,
			&post.UserName,
		)
		totalRecords = post.TotalRecords
		if err != nil {
			return posts, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return posts, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), err
	}
	//log.Println(posts)
	log.Println("pagination data :", utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize))
	return posts, utils.ComputeMetaData(totalRecords, pagenation.Page, pagenation.PageSize), nil

}

func (c *userRepo) VerifyAccount(email string, code int) error {

	var id int

	query := `SELECT id FROM 
				verifications WHERE 
				email = $1 AND code = $2;`
	err := c.db.QueryRow(query, email, code).Scan(&id)

	if err == sql.ErrNoRows {
		return errors.New("Invalid verification code/Email")
	}

	if err != nil {
		return err
	}

	query = `UPDATE users 
				SET
				 verification = $1
				WHERE
				 email = $2 ;`
	err = c.db.QueryRow(query, true, email).Err()
	log.Println("Updating User verification: ", err)
	if err != nil {
		return err
	}

	return nil
}

func (c *userRepo) StoreVerificationDetails(email string, code int) error {

	query := `INSERT INTO 
				verifications(email, code)
				VALUES( $1, $2);`

	err := c.db.QueryRow(query, email, code).Err()

	return err

}

func (c *userRepo) FindUser(email string) (model.UserResponse, error) {

	var user model.UserResponse

	query := `SELECT 
				id,
				email,
				name,
				password_hash
				FROM users 
				WHERE email = $1;`
	query1 := `SELECT id 
				FROM users
				WHERE email = $1 AND activated = true;`
	err := c.db.QueryRow(query1,
		email).Scan(
		&user.ID)
	if err != nil {
		return user, errors.New("blocked")
	}
	err = c.db.QueryRow(query,
		email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
	)

	return user, err
}

func (c *userRepo) InsertUser(user model.User) (int, error) {

	var id int

	query := `INSERT INTO users(
			email,
			name,
			password_hash
			)
			VALUES
			($1, $2, $3)
			RETURNING id;`

	err := c.db.QueryRow(query,
		user.Email,
		user.Name,
		user.Password).Scan(
		&id,
	)
	return id, err
}
