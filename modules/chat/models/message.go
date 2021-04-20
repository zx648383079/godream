package models

import (
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
)

type MessageSimple struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Type      uint32 `json:"type"`
	Content   string `json:"content"`
	CreatedAt uint   `json:"created_at"`
}

type MssageForm struct {
	ItemType    uint   `json:"type"`
	ItemId      uint   `json:"id"`
	MessageType uint   `json:"item_type"`
	Content     string `json:"content"`
	StartTime   uint   `json:"start_time"`
}

type MessageItem struct {
	entities.Message
	User *auth_models.UserSimple
}

func (m MessageItem) IsRoom(itemType uint32, itemId uint) bool {
	if itemId < 1 {
		return false
	}
	if m.GroupId > 0 {
		return itemType > 0 && m.GroupId == itemId
	}
	return itemType < 1 && m.ReceiveId == itemId
}
