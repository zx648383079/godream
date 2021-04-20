package entities

// FriendLink 友情链接Model
type FriendLink struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Logo      string `json:"logo"`
	Brief     string `json:"brief"`
	Email     string `json:"email"`
	Status    int32  `json:"status"`
	UserId    int    `json:"user_id"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (FriendLink) TableName() string {
	return "cif_friend_link"
}
