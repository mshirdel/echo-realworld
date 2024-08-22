package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Uesrname  string         `gorm:"column:username;type:VARCHAR(64);index;unique"`
	Email     string         `gorm:"column:email;type:VARCHAR(256);index;unique"`
	Password  string         `gorm:"column:password;type:VARCHAR(512)"`
	LastLogin *time.Time     `gorm:"column:last_login;type:DATETIME(3)"`
	IsActive  bool           `gorm:"column:is_active;type:BOOLEAN;DEFAULT:true"`
	CreatedAt time.Time      `gorm:"column:created_at;type:DATETIME(3)"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:DATETIME(3)"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (u *User) Run () error {
	if u.ID == 0 {
		return fmt.Errorf("id is zero error: id: [%d]", u.ID)
	}
	
	return nil
}
