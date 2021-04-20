package entities

// GroupUser Model
type GroupUser struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	GroupId   int    `json:"group_id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	RoleId    int    `json:"role_id"`
	Status    int32  `json:"status"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (GroupUser) TableName() string {
	return "chat_group_user"
}
