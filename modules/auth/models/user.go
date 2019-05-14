package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Sex       int
	Avatar    string
	Money     int
	Token     string
	Status    int
	CreatedAt int
	UpdatedAt int
}

func LoginEmail(email string, password string) (user User, err error) {
	err = nil
	if !govalidator.IsEmail(email) {
		err = errors.New("error email")
		return
	}

	return
}
