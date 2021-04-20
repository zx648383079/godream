package entities

// Emoji 表情Model
type Emoji struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	CatId   uint   `json:"cat_id"`
	Name    string `json:"name"`
	Type    uint32 `json:"type"`
	Content string `json:"content"`
}

// TableName 表名
func (Emoji) TableName() string {
	return "seo_emoji"
}

// EmojiCategory 表情分类Model
type EmojiCategory struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// TableName 表名
func (EmojiCategory) TableName() string {
	return "seo_emoji_category"
}
