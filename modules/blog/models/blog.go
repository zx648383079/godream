package models

import (
	"zodream/modules/auth/models"
	"zodream/database"
)

// Blog 博客Model
type Blog struct {
	ID            uint
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

func (Blog) TableName() string {
	return "blog"
}

func GetBlogList() (data []Blog) {
	// var users []models.User
	// database.DB.Where("name = ?", "jinzhu").Find(&data)
	database.DB.Limit(20).Offset(0).Find(&data)
	// database.DB.Model(&data).Related(&users)
	return 
}
