package models

import (
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
)

type ApplyItem struct {
	entities.Apply
	User *models.UserSimple `json:"user"`
}
