package entities

// Group Model
type Group struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	UpdatedAt   uint   `json:"updated_at"`
	CreatedAt   uint   `json:"created_at"`
}

// TableName 表名
func (Group) TableName() string {
	return "chat_group"
}
