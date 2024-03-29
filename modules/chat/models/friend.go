package models

import (
	auth_models "zodream.cn/godream/modules/auth/models"
)

type FriendGroup struct {
	ID     uint      `json:"id"`
	Name   string    `json:"name"`
	Count  uint      `json:"count"`
	Online uint      `json:"online"`
	Users  []*Friend `json:"users"`
}

type Friend struct {
	ID         uint                    `gorm:"primary_key" json:"id"`
	Name       string                  `json:"name"`
	ClassifyId uint                    `json:"classify_id"`
	UserId     uint                    `json:"user_id"`
	Status     uint32                  `json:"status"`
	CreatedAt  uint                    `json:"created_at"`
	User       *auth_models.UserSimple `json:"user"`
}

type FriendSample struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
