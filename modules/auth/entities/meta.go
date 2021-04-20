package entities

// UserMeta Model
type UserMeta struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	UserId  int    `json:"user_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// TableName 表名
func (UserMeta) TableName() string {
	return "user_meta"
}
