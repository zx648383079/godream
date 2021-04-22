package models

// Column 表的列
type Column struct {
	Name     string `gorm:"column:Field"`
	Comment  string `gorm:"column:Comment"`
	DataType string `gorm:"column:Type"`
	Nullable string `gorm:"column:Null"`
	Key      string `gorm:"column:Key"`
	Extra    string `gorm:"column:Extra"`
}
