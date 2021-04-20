package entities

// Group Model
type Group struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
	UpdatedAt   int    `json:"updated_at"`
	CreatedAt   int    `json:"created_at"`
}

// TableName 表名
func (Group) TableName() string {
	return "chat_group"
}
