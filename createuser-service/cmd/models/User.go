package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"-"`
	Password string `json:"-"`
}

func (u *User) Register(DB *gorm.DB) error {
	result := DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return
}
