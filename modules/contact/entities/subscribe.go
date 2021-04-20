package entities

// Subscribe 邮箱订阅Model
type Subscribe struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Email     string `json:"email"`
	Status    int32  `json:"status"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (Subscribe) TableName() string {
	return "cif_subscribe"
}
