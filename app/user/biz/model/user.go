package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Base
	Email          string `gorm:"unique"`
	PasswordHashed string
}

func (u User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}
