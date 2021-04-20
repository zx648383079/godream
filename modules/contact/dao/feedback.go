package dao

import (
	"errors"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/contact/entities"
	"zodream.cn/godream/modules/contact/models"
	"zodream.cn/godream/utils"
)

func SaveFeedback(data *models.FeedbackForm, user int) error {
	if data.Content == "" {
		return errors.New("反馈内容不能为空")
	}
	now := int(utils.Now())
	model := entities.Feedback{
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
		Content:   data.Content,
		UserId:    user,
		Status:    0,
		UpdatedAt: now,
		CreatedAt: now,
	}
	database.DB.Create(&model)
	return nil
}
