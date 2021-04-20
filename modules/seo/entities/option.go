package entities

// Option 全局设置Model
type Option struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	ParentId     int    `json:"parent_id"`
	Type         string `json:"type"`
	Visibility   int32  `json:"visibility"`
	DefaultValue string `json:"default_value"`
	Value        string `json:"value"`
	Position     int32  `json:"position"`
}

// TableName 表名
func (Option) TableName() string {
	return "seo_option"
}
