package service

import (
	"github.com/mshirdel/echo-realworld/app/repo"
	"github.com/mshirdel/echo-realworld/config"
)

type Service struct {
	cfg  *config.Config
	repo *repo.Repository
	User *UserService
}

func New(cfg *config.Config, r *repo.Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: r,
	}
}

func (s *Service) InitAll() error {
	s.User = NewUserService(s.repo)
	return nil
}
