package entities

// Apply Model
type Apply struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	ItemType  uint32 `json:"item_type"`
	ItemId    uint   `json:"item_id"`
	Remark    string `json:"remark"`
	UserId    uint   `json:"user_id"`
	Status    uint32 `json:"status"`
	UpdatedAt uint   `json:"updated_at"`
	CreatedAt uint   `json:"created_at"`
}

// TableName 表名
func (Apply) TableName() string {
	return "chat_apply"
}
