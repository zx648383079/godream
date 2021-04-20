package entities

// Blog 博客Model
type Blog struct {
	ID                  uint   `gorm:"primary_key" json:"id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	Keywords            string `json:"keywords"`
	ParentId            uint   `json:"parent_id"`
	ProgrammingLanguage string `json:"programming_language"`
	Language            string `json:"language"`
	Thumb               string `json:"thumb"`
	EditType            uint32 `json:"edit_type"`
	Content             string `json:"content"`
	UserId              uint   `json:"user_id"`
	TermId              uint   `json:"term_id"`
	Type                uint32 `json:"type"`
	RecommendCount      uint   `json:"recommend_count"`
	CommentCount        uint   `json:"comment_count"`
	ClickCount          uint   `json:"click_count"`
	OpenType            uint32 `json:"open_type"`
	OpenRule            string `json:"open_rule"`
	DeletedAt           uint   `json:"deleted_at"`
	UpdatedAt           uint   `json:"updated_at"`
	CreatedAt           uint   `json:"created_at"`
}

// TableName 表名
func (Blog) TableName() string {
	return "blog"
}
