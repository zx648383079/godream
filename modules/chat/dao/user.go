package dao

import (
	"zodream.cn/godream/database"
	auth_dao "zodream.cn/godream/modules/auth/dao"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
)

func GetProfile(user uint) *models.Profile {
	userModel := auth_dao.GetProfile(user)
	model := models.Profile{
		User: userModel,
		Name: userModel.Name,
	}
	var count int64
	database.DB.Model(&entities.Message{}).Where("receive_id=?", user).Where("status=0").Count(&count)
	model.NewCount = uint(count)
	return &model
}
