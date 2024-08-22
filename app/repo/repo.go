package repo

import (
	"github.com/mshirdel/echo-realworld/app/db"
	"github.com/mshirdel/echo-realworld/config"
)

type Repository struct {
	cfg  *config.Config
	db   *db.DB
	User *UserRepo
}

func New(cfg *config.Config, db *db.DB) *Repository {
	return &Repository{
		cfg: cfg,
		db:  db,
	}
}

func (r *Repository) InitAll() {
	database := r.db.Realworld

	r.User = NewUserRepo(database)
}
