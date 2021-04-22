package models

import "zodream.cn/godream/modules/chat/entities"

type Group struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}

type GroupUserItem struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Sex    uint32 `json:"sex"`
	Avatar string `json:"avatar"`
	Role   uint   `json:"role"`
	Status uint32 `json:"status"`
}

type GroupFull struct {
	entities.Group
	Users []*GroupUserItem
}
