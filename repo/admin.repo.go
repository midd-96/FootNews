package repo

import (
	"NewsAppApi/model"
	"NewsAppApi/utils"
	"database/sql"
	"errors"
	"log"
)

type AdminRepository interface {
	FindAdmin(username string) (model.AdminResponse, error)
	CreateAdmin(admin model.Admin) error
	ApprovePost(post_id int) error
	ListAllPosts(pagenation utils.Filter, sortby string) ([]model.PostResponse, utils.Metadata, error)
}

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}
}

func (c *adminRepo) ListAllPosts(pagenation utils.Filter, sortby string) ([]model.PostResponse, utils.Metadata, error) {

	var posts []model.PostResponse
	log.Println("Pagenation :", pagenation)
	query := `
	SELECT COUNT(*) OVER() AS total_records, pq.*, u.name as uname FROM (
	    SELECT p.id, p.title, p.url, p.created_at, p.user_id as uid, COUNT(c.post_id) as comment_count, count(v.post_id) as votes, p.approved
		FROM posts p 
		LEFT JOIN comments c ON p.id = c.post_id 
	    LEFT JOIN votes v ON p.id = v.post_id
		
		GROUP BY p.id
		ORDER BY p.approved 

		) AS pq
	LEFT JOIN users u ON u.id = uid `

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
		log.Println("Votes : ", post.Votes)
		log.Println(posts)
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

func (c *adminRepo) ApprovePost(post_id int) error {
	query := `UPDATE posts
				SET approved = true
					WHERE id = $1;`
	err := c.db.QueryRow(query, post_id)
	log.Println(err)
	if err == nil {
		return errors.New("Error while approving the post")
	}

	return nil
}

func (c *adminRepo) FindAdmin(email string) (model.AdminResponse, error) {

	var admin model.AdminResponse

	query := `SELECT
			id, 
			email,
			password_hash
			FROM admin WHERE email = $1;`

	err := c.db.QueryRow(query,
		email).Scan(
		&admin.ID,
		&admin.Email,
		&admin.Password)

	return admin, err
}

func (c *adminRepo) CreateAdmin(admin model.Admin) error {

	query := `INSERT INTO
				admin (email,name,password_hash)
				VALUES
				(
					$1, $2, $3
					);`
	err := c.db.QueryRow(
		query, admin.Email,
		admin.Name,
		admin.Password,
	).Err()
	return err
}
