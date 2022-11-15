package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID              int    `json:"id" gorm:"primary_key"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"hash_password"`
	ConfirmPassword string `json:"confirm_password"`
}

// user schema for user table
type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"hash_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Post struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	UserID    int       `json:"user_id"`
}

type Comment struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
}

type Vote struct {
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
}

type Report struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	PostID    int       `json:"post_id"`
}

//to store mail verification details

type Verification struct {
	gorm.Model

	Email string `json:"email"`
	Code  int    `json:"code"`
}
