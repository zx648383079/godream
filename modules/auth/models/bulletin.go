package models

type Bulletin struct {
	ID        uint
	Title     string
	Content   string
	Type      int
	UserID    int
	CreatedAt int
	UpdatedAt int
}
