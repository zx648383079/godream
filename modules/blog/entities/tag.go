package entities

// Tag Model
type Tag struct {
	ID          uint `gorm:"primary_key" json:"id"`
	Name        string
	Description string
	BlogCount   int
}

// TableName 表名
func (Tag) TableName() string {
	return "blog_tag"
}

// TagRelationship Model
type TagRelationship struct {
	TagId    int
	BlogId   int
	Position int32
}

// TableName 表名
func (TagRelationship) TableName() string {
	return "blog_tag_relationship"
}
