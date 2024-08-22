package service

import (
	"context"
	"fmt"

	"github.com/mshirdel/echo-realworld/app/repo"
	"github.com/mshirdel/echo-realworld/models"
)

type UserService struct {
	repo *repo.Repository
}

func NewUserService(r *repo.Repository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (u *UserService) RegisterUser(ctx context.Context, user models.User) error {
	err := u.repo.User.Create(ctx, user)
	// todo log error
	if err != nil {
		return fmt.Errorf("error in creating user")
	}

	return nil
}

func (u *UserService) LoginUser(ctx context.Context, username string, password string) {
	// repo.user.creat()
}
