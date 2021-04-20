package entities

// Bulletin Model
type Bulletin struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ExtraRule string `json:"extra_rule"`
	Type      uint32 `json:"type"`
	UserId    uint   `json:"user_id"`
	UpdatedAt uint   `json:"updated_at"`
	CreatedAt uint   `json:"created_at"`
}

// TableName 表名
func (Bulletin) TableName() string {
	return "bulletin"
}

// BulletinUser Model
type BulletinUser struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	BulletinId uint   `json:"bulletin_id"`
	Status     uint32 `json:"status"`
	UserId     uint   `json:"user_id"`
	UpdatedAt  uint   `json:"updated_at"`
	CreatedAt  uint   `json:"created_at"`
}

// TableName 表名
func (BulletinUser) TableName() string {
	return "bulletin_user"
}
