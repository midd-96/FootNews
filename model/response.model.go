package model

import "time"

type PostResponse struct {
	ID           int       `json:"id,omitempty"`
	Title        string    `json:"title"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       int       `json:"user_id"`
	Votes        int       `json:"votes,omitempty"`
	UserName     string    `json:"user_name,omitempty"`
	CommentCount int       `json:"comment_count,omitempty"`
	Approved     bool      `json:"approved"`
	TotalRecords int       `json:"total_records,omitempty"`
}

type AdminResponse struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Created_At time.Time `json:"created_at"`
	Email      string    `json:"email"`
	//Name       string    `json:"name"`
	Password string `json:"hash_password"`
	Token    string `json:"token,omitempty"`
}

type UserResponse struct {
	ID         int       `json:"id"`
	Created_At time.Time `json:"created_at"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Password   string    `json:"password"`
	Token      string    `json:"token"`
}
