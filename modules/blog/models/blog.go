package models

import (
	"log"
	"zodream/database"
	"zodream/modules/auth/models"
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

func (Blog) TableName() string {
	return "blog"
}

func GetBlogList() (data []Blog) {
	//var users []models.User
	// database.DB.Where("name = ?", "jinzhu").Find(&data)
	database.DB.LogMode(true)
	database.DB.Unscoped().Where("deleted_at=0").Preload("User").Limit(2).Offset(0).Find(&data)
	log.Println(data[0].User.Name)
	return
}
