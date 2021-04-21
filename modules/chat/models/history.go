package models

import (
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
)

type HistoryItem struct {
	entities.History
	User    *auth_models.UserSimple `json:"user"`
	Friend  *FriendSample           `json:"friend"`
	Group   *Group                  `json:"group"`
	Message *MessageSimple          `json:"message"`
}
