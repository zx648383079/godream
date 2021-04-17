package entities

// BlogMeta Model
type BlogMeta struct {
	ID      uint `gorm:"primary_key" json:"id"`
	BlogId  int
	Name    string
	Content string
}

// TableName 表名
func (BlogMeta) TableName() string {
	return "blog_meta"
}
