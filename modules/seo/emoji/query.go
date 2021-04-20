package emoji

import (
	"zodream.cn/godream/database"
	"zodream.cn/godream/modules/seo/entities"
	"zodream.cn/godream/modules/seo/models"
)

func GetAll() []*models.EmojiCategoryItem {
	var categories []*entities.EmojiCategory
	database.DB.Find(&categories)
	var l = len(categories)
	items := make([]*models.EmojiCategoryItem, l)
	if l < 1 {
		return items
	}
	var data []*entities.Emoji
	database.DB.Find(&data)
	for i := 0; i < l; i++ {
		v := categories[i]
		item := models.EmojiCategoryItem{
			EmojiCategory: *v,
		}
		for _, e := range data {
			if v.ID == e.CatId {
				item.Items = append(item.Items, e)
			}
		}
		items[i] = &item
	}
	return items
}
