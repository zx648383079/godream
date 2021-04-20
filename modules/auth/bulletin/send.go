package bulletin

import (
	"encoding/json"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/auth/entities"
	"zodream.cn/godream/utils"
	"zodream.cn/godream/utils/rule"
)

const (
	BULLETIN_TYPE = 0
)

func SendSystem(users []uint, title string, content string, bulletinType uint32, extraRule []rule.RuleItem) {
	Send(0, users, title, content, bulletinType, extraRule)
}

func Send(sender uint, users []uint, title string, content string, bulletinType uint32, extraRule []rule.RuleItem) {
	ruleByte, _ := json.Marshal(extraRule)
	now := uint(utils.Now())
	model := entities.Bulletin{
		Title:     title,
		Content:   content,
		Type:      bulletinType,
		UserId:    sender,
		ExtraRule: string(ruleByte),
		CreatedAt: now,
		UpdatedAt: now,
	}
	smt := database.DB.Create(&model)
	if smt.Error != nil {
		return
	}
	var items []entities.BulletinUser
	for _, u := range users {
		exist := false
		for _, v := range items {
			if v.UserId == u {
				exist = true
				break
			}
		}
		if exist {
			continue
		}
		items = append(items, entities.BulletinUser{
			BulletinId: model.ID,
			UserId:     u,
			Status:     0,
			CreatedAt:  now,
			UpdatedAt:  now,
		})
	}
	database.DB.CreateInBatches(items, 100)
}
