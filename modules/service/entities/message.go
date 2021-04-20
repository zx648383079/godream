package entities

// Message Model
type Message struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserId    int    `json:"user_id"`
	SessionId int    `json:"session_id"`
	SendType  int32  `json:"send_type"`
	Type      int32  `json:"type"`
	Content   string `json:"content"`
	Status    int32  `json:"status"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (Message) TableName() string {
	return "service_message"
}
