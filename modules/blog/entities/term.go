package entities

// Term Model
type Term struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	ParentId    uint   `json:"parent_id"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
	Styles      string `json:"styles"`
}

// TableName 表名
func (Term) TableName() string {
	return "blog_term"
}
