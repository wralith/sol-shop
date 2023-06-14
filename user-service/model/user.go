package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password []byte

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CreateUserOptions struct {
	Username string
	Email    string
	Password []byte
}

func CreateUser(opts CreateUserOptions) User {
	return User{
		Username: opts.Username,
		Email:    opts.Email,
		Password: opts.Password,
	}
}
