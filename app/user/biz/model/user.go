package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Base
	Email          string `gorm:"unique"`
	PasswordHashed string
	Role           uint32
}

func (u User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, ctx context.Context, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(db *gorm.DB, ctx context.Context, email string) (user *User, err error) {
	err = db.WithContext(ctx).Model(&User{}).Where(&User{Email: email}).First(&user).Error
	return
}
