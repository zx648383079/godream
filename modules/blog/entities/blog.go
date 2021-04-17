package entities

// Blog 博客Model
type Blog struct {
	ID                  uint `gorm:"primary_key" json:"id"`
	Title               string
	Description         string
	Keywords            string
	ParentId            int
	ProgrammingLanguage string
	Language            string
	Thumb               string
	EditType            int32
	Content             string
	UserId              int
	TermId              int
	Type                int32
	RecommendCount      int
	CommentCount        int
	ClickCount          int
	OpenType            int32
	OpenRule            string
	DeletedAt           int
	UpdatedAt           int
	CreatedAt           int
}

// TableName 表名
func (Blog) TableName() string {
	return "blog"
}
