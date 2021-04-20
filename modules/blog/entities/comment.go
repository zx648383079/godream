package entities

// Comment Model
type Comment struct {
	ID            uint   `gorm:"primary_key" json:"id"`
	Content       string `json:"content"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Url           string `json:"url"`
	ParentId      uint   `json:"parent_id"`
	Position      uint   `json:"position"`
	UserId        uint   `json:"user_id"`
	BlogId        uint   `json:"blog_id"`
	Ip            string `json:"ip"`
	Agent         string `json:"agent"`
	AgreeCount    uint   `json:"agree_count"`
	DisagreeCount uint   `json:"disagree_count"`
	Approved      uint32 `json:"approved"`
	CreatedAt     uint   `json:"created_at"`
}

// TableName 表名
func (Comment) TableName() string {
	return "blog_comment"
}
