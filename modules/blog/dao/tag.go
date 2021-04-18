package dao

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/entities"
)

func GetTags() []*entities.Tag {
	var items []*entities.Tag
	database.DB.Find(&items)
	return items
}
