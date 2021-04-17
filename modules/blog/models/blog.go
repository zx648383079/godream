package models

import (
	"zodream.cn/godream/modules/auth/models"
	"zodream.cn/godream/modules/blog/entities"
)

// BlogPageItem 博客Model
type BlogPageItem struct {
	ID             uint `gorm:"primary_key" json:"id"`
	Title          string
	Description    string
	Keywords       string
	Language       string
	Thumb          string
	User           models.UserSimple
	UserID         int
	Term           entities.Term
	TermID         int
	Type           int
	RecommendCount int
	CommentCount   int
	ClickCount     int
	OpenType       int32
	UpdatedAt      int
	CreatedAt      int
}

// TableName 表名
func (BlogPageItem) TableName() string {
	return "blog"
}
