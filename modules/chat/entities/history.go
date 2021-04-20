package entities

// History Model
type History struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	ItemType    uint32 `json:"item_type"`
	ItemId      uint   `json:"item_id"`
	UserId      uint   `json:"user_id"`
	UnreadCount uint   `json:"unread_count"`
	LastMessage uint   `json:"last_message"`
	UpdatedAt   uint   `json:"updated_at"`
	CreatedAt   uint   `json:"created_at"`
}

// TableName 表名
func (History) TableName() string {
	return "chat_history"
}
