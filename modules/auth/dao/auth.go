package dao

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/auth/entities"
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/utils"
)

func LoginEmail(data models.LoginEmail) (*entities.User, error) {
	if !govalidator.IsEmail(data.Email) {
		err := errors.New("error email")
		return nil, err
	}
	var user entities.User
	err := database.DB.Where("email=?", data.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	if !utils.PasswordVerify(data.Password, user.Password) {
		err = errors.New("error password")
		return nil, err
	}
	return &user, err
}
