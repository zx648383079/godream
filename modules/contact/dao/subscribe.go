package dao

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/contact/entities"
	"zodream.cn/godream/utils"
)

func Subscribe(email string) error {
	if govalidator.IsEmail(email) {
		return errors.New("邮箱错误")
	}
	now := int(utils.Now())
	model := entities.Subscribe{
		Email:     email,
		Status:    0,
		UpdatedAt: now,
		CreatedAt: now,
	}
	var count int64
	database.DB.Model(&model).Where("email=?", email).Count(&count)
	if count > 0 {
		return errors.New("邮箱已订阅过")
	}
	database.DB.Create(&model)
	return nil
}

func Unsubscribe(email string) error {
	if govalidator.IsEmail(email) {
		return errors.New("邮箱错误")
	}
	database.DB.Where("email=?", email).Delete(&entities.Subscribe{})
	return nil
}
