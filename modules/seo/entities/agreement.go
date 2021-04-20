package entities

// Agreement 服务协议Model
type Agreement struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Status      int32  `json:"status"`
	UpdatedAt   int    `json:"updated_at"`
	CreatedAt   int    `json:"created_at"`
}

// TableName 表名
func (Agreement) TableName() string {
	return "seo_agreement"
}
