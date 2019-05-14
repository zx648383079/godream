package models

type BookChapter struct {
	ID        uint
	BookID    int
	Title     string
	ParentID  int
	Status    int
	Position  int
	Size      int
	Source    string
	DeletedAt int
	CreatedAt int
	UpdatedAt int
}
