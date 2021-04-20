package entities

// Message Model
type Message struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Type      uint32 `json:"type"`
	Content   string `json:"content"`
	ExtraRule string `json:"extra_rule"`
	ItemId    uint   `json:"item_id"`
	ReceiveId uint   `json:"receive_id"`
	GroupId   uint   `json:"group_id"`
	UserId    uint   `json:"user_id"`
	Status    uint32 `json:"status"`
	DeletedAt uint   `json:"deleted_at"`
	UpdatedAt uint   `json:"updated_at"`
	CreatedAt uint   `json:"created_at"`
}

// TableName 表名
func (Message) TableName() string {
	return "chat_message"
}
