package models

import (
	"zodream.cn/godream/modules/auth/models"
)

// Blog 博客Model
type Blog struct {
	ID            uint `gorm:"primary_key" json:"id"`
	Title         string
	Description   string
	Keywords      string
	Language      string
	Thumb         string
	EditType      int
	Content       string
	User          models.User
	UserID        int
	Term          BlogTerm
	TermID        int
	Type          int
	SourceURL     string
	Recommend     int
	CommentCount  int
	ClickCount    int
	CommentStatus int
	DeletedAt     int
	CreatedAt     int
	UpdatedAt     int
}

// TableName 表名
func (Blog) TableName() string {
	return "blog"
}
