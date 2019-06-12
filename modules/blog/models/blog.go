package models

type Blog struct {
	ID            uint
	Title         string
	Description   string
	Keywords      string
	Language      string
	Thumb         string
	EditType      int
	Content       string
	UserID        int
	TermID        int
	Type          int
	SourceUrl     string
	Recommend     int
	CommentCount  int
	ClickCount    int
	CommentStatus int
	DeletedAt     int
	CreatedAt     int
	UpdatedAt     int
}
