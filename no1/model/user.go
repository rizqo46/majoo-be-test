package model

import (
	"crypto/md5"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	UserName  string    `json:"user_name" gorm:"column:user_name"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	CreatedBy string    `json:"created_by" gorm:"column:created_by"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	UpdatedBy string    `json:"updated_by" gorm:"column:updated_by"`
}

func (t *User) TableName() string {
	return "Users"
}

func (t *User) EncryptPassword() {
	t.Password = encryptPassword(t.Password)
}

func encryptPassword(pass string) string {
	data := []byte(pass)
	return fmt.Sprintf("%x", md5.Sum(data))
}
