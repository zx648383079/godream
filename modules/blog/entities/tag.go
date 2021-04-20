package entities

// Tag Model
type Tag struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BlogCount   uint   `json:"blog_count"`
}

// TableName 表名
func (Tag) TableName() string {
	return "blog_tag"
}

// TagRelationship Model
type TagRelationship struct {
	TagId    uint
	BlogId   uint
	Position uint32
}

// TableName 表名
func (TagRelationship) TableName() string {
	return "blog_tag_relationship"
}
