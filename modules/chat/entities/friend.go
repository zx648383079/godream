package entities

// Friend Model
type Friend struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	ClassifyId uint   `json:"classify_id"`
	UserId     uint   `json:"user_id"`
	BelongId   uint   `json:"belong_id"`
	Status     uint32 `json:"status"`
	UpdatedAt  uint   `json:"updated_at"`
	CreatedAt  uint   `json:"created_at"`
}

// TableName 表名
func (Friend) TableName() string {
	return "chat_friend"
}
