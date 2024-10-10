package repo

import (
	"context"

	"github.com/mshirdel/echo-realworld/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error

	return user, err
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Find(&user, id).Error

	return user, err
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *UserRepo) Update(ctx context.Context, user models.User) error {
	return r.db.Save(&user).Error
}

func (r *UserRepo) Find(ctx context.Context, users *[]models.User) error {
	return r.db.WithContext(ctx).Select("username").Find(users).Error
}
