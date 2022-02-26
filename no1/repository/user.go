package repository

import (
	"majoo-be-test/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Login(user *model.User) error
}

type userRepoImpl struct {
	dB *gorm.DB
}

func NewUserRepo(dB *gorm.DB) UserRepo {
	return &userRepoImpl{dB}
}

func (r *userRepoImpl) Login(user *model.User) error {
	user.EncryptPassword()
	if err := r.dB.Where(&user).First(&user).Error; err != nil {
		return err
	}
	return nil
}
