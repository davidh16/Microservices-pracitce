package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"-"`
	Password string `json:"-"`
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
