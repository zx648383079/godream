package models

type Book struct {
	ID             uint
	Name           string
	Cover          string
	Description    string
	AuthorID       int
	UserID         int
	Classify       int
	CatID          int
	Size           int
	ClickCount     int
	RecommendCount int
	OverAt         int
	Source         string
	DeletedAt      int
	CreatedAt      int
	UpdatedAt      int
}
