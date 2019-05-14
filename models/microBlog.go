package models

type MicroBlog struct {
	ID        uint
	UserID    int
	Content   string
	Recommend int
	Source    string
	CreatedAt int
	UpdatedAt int
}
