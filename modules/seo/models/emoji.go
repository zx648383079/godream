package models

import "zodream.cn/godream/modules/seo/entities"

type EmojiCategoryItem struct {
	entities.EmojiCategory
	Items []*entities.Emoji
}
