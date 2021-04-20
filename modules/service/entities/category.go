package entities

// Category Model
type Category struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

// TableName 表名
func (Category) TableName() string {
	return "service_category"
}

// CategoryUser Model
type CategoryUser struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CatId     int  `json:"cat_id"`
	UserId    int  `json:"user_id"`
	UpdatedAt int  `json:"updated_at"`
	CreatedAt int  `json:"created_at"`
}

// TableName 表名
func (CategoryUser) TableName() string {
	return "service_category_user"
}

// CategoryWord Model
type CategoryWord struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Content string `json:"content"`
	CatId   int    `json:"cat_id"`
}

// TableName 表名
func (CategoryWord) TableName() string {
	return "service_category_word"
}
