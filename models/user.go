package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"column:username;type:VARCHAR(64);index;unique"`
	Email     string         `gorm:"column:email;type:VARCHAR(256);index;unique"`
	Password  string         `gorm:"column:password;type:VARCHAR(512)"`
	LastLogin time.Time      `gorm:"column:last_login;type:TIMESTAMP;null;default:null"`
	IsActive  bool           `gorm:"column:is_active;type:BOOLEAN;DEFAULT:true"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
