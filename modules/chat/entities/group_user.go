package entities

// GroupUser Model
type GroupUser struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	GroupId   uint   `json:"group_id"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	RoleId    uint   `json:"role_id"`
	Status    uint32 `json:"status"`
	UpdatedAt uint   `json:"updated_at"`
	CreatedAt uint   `json:"created_at"`
}

// TableName 表名
func (GroupUser) TableName() string {
	return "chat_group_user"
}
