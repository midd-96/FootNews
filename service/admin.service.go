package service

import (
	"NewsAppApi/model"
	"NewsAppApi/repo"
	"NewsAppApi/utils"
	"database/sql"
	"errors"
	"log"
)

type AdminService interface {
	FindAdmin(username string) (*model.AdminResponse, error)
	CreateAdmin(admin model.Admin) error
	ApprovePost(post_id int) error
	ListAllPosts(pagenation utils.Filter, sortby string) (*[]model.PostResponse, *utils.Metadata, error)
}

type adminService struct {
	adminRepo repo.AdminRepository
	userRepo  repo.UserRepository
}

func NewAdminService(
	adminRepo repo.AdminRepository,
	userRepo repo.UserRepository,
) AdminService {
	return &adminService{
		adminRepo: adminRepo,
		userRepo:  userRepo,
	}
}

func (c *adminService) ListAllPosts(pagenation utils.Filter, sortby string) (*[]model.PostResponse, *utils.Metadata, error) {

	posts, metadata, err := c.adminRepo.ListAllPosts(pagenation, sortby)
	log.Println("metadata from service", metadata)
	if err != nil {
		return nil, &metadata, err
	}

	return &posts, &metadata, nil
}

func (c *adminService) ApprovePost(post_id int) error {
	err := c.adminRepo.ApprovePost(post_id)
	if err != nil {
		return errors.New("Error while approving the post")
	}
	return nil
}

func (c *adminService) CreateAdmin(admin model.Admin) error {

	_, err := c.adminRepo.FindAdmin(admin.Email)

	if err == nil {
		return errors.New("Admin already exists")
	}

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//hashing password
	admin.Password = HashPassword(admin.Password)
	err = c.adminRepo.CreateAdmin(admin)

	if err != nil {
		return errors.New("error while signup")
	}
	return nil
}

func (c *adminService) FindAdmin(email string) (*model.AdminResponse, error) {
	admin, err := c.adminRepo.FindAdmin(email)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
