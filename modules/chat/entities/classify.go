package entities

// FriendClassify Model
type FriendClassify struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	UserId    uint   `json:"user_id"`
	CreatedAt uint   `json:"created_at"`
}

// TableName 表名
func (FriendClassify) TableName() string {
	return "chat_friend_classify"
}
