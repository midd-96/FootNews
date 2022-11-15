package service

import (
	"NewsAppApi/config"
	"NewsAppApi/model"
	"NewsAppApi/repo"
	"NewsAppApi/utils"
	"crypto/md5"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type UserService interface {
	ReportPost(report model.Report) error
	MarkVote(newVote model.Vote) (error, string)
	AddComment(newComment model.Comment) error
	AllPosts(pagenation utils.Filter, sortby string) (*[]model.PostResponse, *utils.Metadata, error)
	FindUser(email string) (*model.UserResponse, error)
	CreateUser(newUser model.User) error
	AddPost(newPost model.Post) (int, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
}

type userService struct {
	userRepo   repo.UserRepository
	adminRepo  repo.AdminRepository
	mailConfig config.MailConfig
}

func NewUserService(
	userRepo repo.UserRepository,
	adminRepo repo.AdminRepository,
	mailConfig config.MailConfig) UserService {
	return &userService{
		userRepo:   userRepo,
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
	}
}

func (c *userService) ReportPost(report model.Report) error {
	err := c.userRepo.ReportPost(report)
	if err != nil {
		return err
	}
	return nil
}

func (c *userService) MarkVote(newVote model.Vote) (error, string) {
	err, msg := c.userRepo.MarkVote(newVote)
	if err != nil {
		return err, msg
	}
	return nil, msg
}

func (c *userService) AddComment(newComment model.Comment) error {
	err := c.userRepo.AddComment(newComment)
	if err != nil {
		return err
	}
	return nil
}

func (c *userService) AddPost(newPost model.Post) (int, error) {
	_, err := c.userRepo.AddPost(newPost)
	if err != nil {
		return newPost.ID, err
	}
	return newPost.ID, nil
}

func (c *userService) AllPosts(pagenation utils.Filter, sortby string) (*[]model.PostResponse, *utils.Metadata, error) {

	posts, metadata, err := c.userRepo.AllPosts(pagenation, sortby)
	log.Println("metadata from service", metadata)
	if err != nil {
		return nil, &metadata, err
	}

	return &posts, &metadata, nil
}

func (c *userService) CreateUser(newUser model.User) error {

	_, err := c.userRepo.FindUser(newUser.Email)

	if err == nil {
		return errors.New("Username already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	newUser.Password = HashPassword(newUser.Password)

	_, err = c.userRepo.InsertUser(newUser)
	if err != nil {
		return err
	}
	return nil

}

func (c *userService) VerifyAccount(email string, code int) error {

	err := c.userRepo.VerifyAccount(email, code)

	if err != nil {
		return err
	}
	return nil
}

func (c *userService) SendVerificationEmail(email string) error {
	//to generate random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(100000)

	message := fmt.Sprintf(
		"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for using NewsAppApi.\n with regards Team NewsAppApi.",
		code,
	)

	// send random code to user's email
	if err := c.mailConfig.SendMail(email, message); err != nil {
		return err
	}

	err := c.userRepo.StoreVerificationDetails(email, code)

	if err != nil {
		return err
	}

	return nil
}

func (c *userService) FindUser(email string) (*model.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password

}
