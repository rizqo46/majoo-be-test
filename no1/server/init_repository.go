package server

import (
	"majoo-be-test/repository"

	"gorm.io/gorm"
)

type Repository struct {
	UserRepo     repository.UserRepo
	MerchantRepo repository.MerchantRepo
}

func InitRepository(dB *gorm.DB) Repository {
	return Repository{
		UserRepo:     repository.NewUserRepo(dB),
		MerchantRepo: repository.NewMerchantRepo(dB),
	}
}
