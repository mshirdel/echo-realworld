package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mshirdel/echo-realworld/app/repo"
	"github.com/mshirdel/echo-realworld/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repo.Repository
}

func NewUserService(r *repo.Repository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (u *UserService) RegisterUser(ctx context.Context, user models.User) (uint, error) {
	pass, err := makePassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("error in encrypting password: %w", err)
	}

	user.Password = pass

	user, err = u.repo.User.Create(ctx, user)
	// todo log error
	if err != nil {
		return 0, fmt.Errorf("error in creating user: %w", err)
	}

	return user.ID, nil
}

func (u *UserService) LoginUser(ctx context.Context, username string, password string) (models.User, error) {
	user, err := u.repo.User.FindByEmail(ctx, username)
	if err != nil {
		return models.User{}, fmt.Errorf("you are not authorized")
	}

	err = checkPassword(password, user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("you are not authorized")
	}

	user.LastLogin = time.Now()
	err = u.repo.User.Update(ctx, user)
	if err != nil {
		// and log this
		return models.User{}, fmt.Errorf("you are not authorized")
	}

	return user, nil
}

func (u *UserService) UserInfo(ctx context.Context, id uint) (models.User, error) {
	return u.repo.User.FindByID(ctx, id)
}

func makePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func checkPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
