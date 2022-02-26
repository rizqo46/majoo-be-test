package database

import "gorm.io/gorm"

func Connect() *gorm.DB {
	return &gorm.DB{}
}
