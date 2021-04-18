package dao

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/blog/entities"
)

func GetCategories() []*entities.Term {
	var items []*entities.Term
	database.DB.Find(&items)
	return items
}
