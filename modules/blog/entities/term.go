package entities

// Term Model
type Term struct {
	ID          uint `gorm:"primary_key" json:"id"`
	Name        string
	ParentId    int
	Keywords    string
	Description string
	Thumb       string
	Styles      string
}

// TableName 表名
func (Term) TableName() string {
	return "blog_term"
}
