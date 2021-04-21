package dao

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/auth/models"
)

func GetProfile(id uint) *models.UserSimple {
	var user models.UserSimple
	database.DB.Select("id,name,email,avatar,sex").First(&user, id)
	if user.ID < 1 {
		return nil
	}
	return &user
}
