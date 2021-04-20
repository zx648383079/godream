package entities

// Feedback 留言Model
type Feedback struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Content   string `json:"content"`
	Status    int32  `json:"status"`
	UserId    int    `json:"user_id"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (Feedback) TableName() string {
	return "cif_feedback"
}
