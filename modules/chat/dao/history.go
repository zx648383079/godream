package dao

import (
	"zodream.cn/godream/database"
	auth_models "zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/chat/entities"
	"zodream.cn/godream/modules/chat/models"
	"zodream.cn/godream/utils"
)

func GetHistories(user int) []*models.HistoryItem {
	var histories []*entities.History
	database.DB.Where("user_id=?", user).Order("updated_at desc").Find(histories)
	items := make([]*models.HistoryItem, len(histories))
	if len(histories) < 1 {
		return items
	}
	var userId []uint
	var groupId []uint
	var messageId []uint
	for _, v := range histories {
		exsit := false
		if v.ItemType < 1 {
			for _, u := range userId {
				if u == v.ItemId {
					exsit = true
					break
				}
			}
			if !exsit {
				userId = append(userId, v.ItemId)
			}
		} else {
			for _, u := range groupId {
				if u == v.ItemId {
					exsit = true
					break
				}
			}
			if !exsit {
				groupId = append(groupId, v.ItemId)
			}
		}
		if v.LastMessage < 1 {
			continue
		}
		exsit = false
		for _, u := range messageId {
			if u == v.LastMessage {
				exsit = true
				break
			}
		}
		if !exsit {
			messageId = append(messageId, v.LastMessage)
		}
	}
	var messageItems []*entities.Message
	if len(messageId) > 0 {
		database.DB.Find(&messageItems, messageId)
	}
	var userItems []*auth_models.UserSimple
	var friendItems []*entities.Friend
	if len(userId) > 0 {
		database.DB.Find(&userItems, userId)
		database.DB.Where("user_id in ?", userId).Where("belong_id=?", user).Find(&friendItems)
	}
	var groupItems []*entities.Group
	if len(groupId) > 0 {
		database.DB.Find(&groupItems, groupId)
	}
	for i := len(histories) - 1; i >= 0; i-- {
		v := histories[i]
		item := models.HistoryItem{
			History: *v,
		}
		if v.ItemType < 1 {
			for _, u := range userItems {
				if u.ID == v.ItemId {
					item.User = u
					break
				}
			}
			for _, f := range friendItems {
				if f.UserId == v.ItemId {
					item.Friend = &models.FriendSample{
						ID:   f.ID,
						Name: f.Name,
					}
					break
				}
			}
		} else {
			for _, g := range groupItems {
				if g.ID == v.ItemId {
					item.Group = &models.Group{
						ID:   g.ID,
						Name: g.Name,
						Logo: g.Logo,
					}
					break
				}
			}
		}
		if v.LastMessage > 0 {
			for _, m := range messageItems {
				if m.ID == v.LastMessage {
					item.Message = &models.MessageSimple{
						ID:        m.ID,
						Type:      m.Type,
						Content:   m.Content,
						CreatedAt: m.CreatedAt,
					}
					break
				}
			}
		}
		items[i] = &item
	}
	return items
}

func addHistory(user uint, itemType uint, itemId uint, messageId uint, count uint) {
	var model entities.History
	database.DB.Where("user_id=?", user).Where("item_type=?", itemType).Where("item_id=?", itemId).First(&model)
	now := uint(utils.Now())
	model.UpdatedAt = now
	model.LastMessage = messageId
	model.UnreadCount += count
	if model.ID > 0 {
		database.DB.Save(&model)
		return
	}
	model.ItemType = uint32(itemType)
	model.ItemId = itemId
	model.CreatedAt = now
	model.UserId = user
	database.DB.Create(&model)
}

func RemoveHistory(user uint, itemType uint, itemId uint) {
	var model entities.History
	database.DB.Where("user_id=?", user).Where("item_type=?", itemType).Where("item_id=?", itemId).Delete(model)
}
