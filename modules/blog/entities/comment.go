package entities

// Comment Model
type Comment struct {
	ID            uint `gorm:"primary_key" json:"id"`
	Content       string
	Name          string
	Email         string
	Url           string
	ParentId      int
	Position      int
	UserId        int
	BlogId        int
	Ip            string
	Agent         string
	AgreeCount    int
	DisagreeCount int
	Approved      int32
	CreatedAt     int
}

// TableName 表名
func (Comment) TableName() string {
	return "blog_comment"
}
