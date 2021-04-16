package models

import (
	"errors"

	"zodream.cn/godream/database"
	"zodream.cn/godream/utils"

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

func (User) TableName() string {
	return "user"
}

func LoginEmail(email string, password string) (*User, error) {
	if !govalidator.IsEmail(email) {
		err := errors.New("error email")
		return nil, err
	}
	var user User
	err := database.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	if !utils.PasswordVerify(password, user.Password) {
		err = errors.New("error password")
		return nil, err
	}
	return &user, err
}
