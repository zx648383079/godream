package dao

import (
	"errors"

	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/contact/entities"
	"zodream.cn/godream/modules/contact/models"
	"zodream.cn/godream/utils"
)

func GetLinkList() []*entities.FriendLink {
	var items []*entities.FriendLink
	database.DB.Where("status=1").Find(&items)
	return items
}

func ApplyLink(data *models.LinkForm, user int) error {
	if data.Name == "" || data.Url == "" {
		return errors.New("站点名或网址不能为空")
	}
	now := int(utils.Now())
	model := entities.FriendLink{
		Name:      data.Name,
		Email:     data.Email,
		Url:       data.Url,
		Brief:     data.Brief,
		Logo:      data.Logo,
		UserId:    user,
		Status:    0,
		UpdatedAt: now,
		CreatedAt: now,
	}
	var count int64
	database.DB.Model(&model).Where("url=?", data.Url).Count(&count)
	if count > 0 {
		return errors.New("网址已申请过了")
	}
	database.DB.Create(&model)
	return nil
}
