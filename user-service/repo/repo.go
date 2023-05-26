package repo

import (
	"errors"

	"github.com/wralith/sol-shop/user-service/model"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(user model.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return errors.Join(errors.New("error while creating user"), err)
	}
	return nil
}

func (r *Repo) FindByID(id uint) (user model.User, err error) {
	err = r.db.First(&user, id).Error
	if err != nil {
		return user, errors.Join(errors.New("error while finding user by id"), err)
	}
	return user, nil
}

func (r *Repo) FindByEmail(email string) (user model.User, err error) {
	err = r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errors.Join(errors.New("error while finding user by email"), err)
	}
	return user, nil
}
