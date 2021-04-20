package models

import (
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/blog/entities"
)

// BlogPageItem 博客Model
type BlogPageItem struct {
	ID                  uint              `gorm:"primary_key" json:"id"`
	Title               string            `json:"title"`
	Description         string            `json:"description"`
	Keywords            string            `json:"keywords"`
	ProgrammingLanguage string            `json:"programming_language"`
	Language            string            `json:"language"`
	Thumb               string            `json:"thumb"`
	User                models.UserSimple `json:"user"`
	UserID              uint              `json:"user_id"`
	Term                entities.Term     `json:"term"`
	TermId              uint              `json:"term_id"`
	Type                uint32            `json:"type"`
	RecommendCount      uint              `json:"recommend_count"`
	CommentCount        uint              `json:"comment_count"`
	ClickCount          uint              `json:"click_count"`
	OpenType            uint32            `json:"open_type"`
	UpdatedAt           uint              `json:"updated_at"`
	CreatedAt           uint              `json:"created_at"`
}

// TableName 表名
func (BlogPageItem) TableName() string {
	return "blog"
}
