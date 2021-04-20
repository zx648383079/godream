package entities

// Apply Model
type Apply struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	ItemType  int32  `json:"item_type"`
	ItemId    int    `json:"item_id"`
	Remark    string `json:"remark"`
	UserId    int    `json:"user_id"`
	Status    int32  `json:"status"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (Apply) TableName() string {
	return "chat_apply"
}
