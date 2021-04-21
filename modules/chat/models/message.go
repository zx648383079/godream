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

type MssageQuery struct {
	ItemType  uint32 `json:"type"`
	ItemId    uint   `json:"id"`
	StartTime uint   `json:"start_time"`
}

type MssageForm struct {
	ItemType    uint   `form:"type" json:"type"`
	ItemId      uint   `form:"id" json:"id"`
	MessageType uint   `json:"item_type"`
	Content     string `json:"content"`
	StartTime   uint   `form:"start_time" json:"start_time"`
	File        string `json:"file"`
	FileName    string `json:"file_name"`
}

type MessageItem struct {
	entities.Message
	User *auth_models.UserSimple `json:"user"`
}

func (m MessageItem) IsRoom(itemType uint32, itemId uint) bool {
	if itemId < 1 {
		return false
	}
	if m.GroupId > 0 {
		return itemType > 0 && m.GroupId == itemId
	}
	return itemType < 1 && (m.ReceiveId == itemId || m.UserId == itemId)
}
