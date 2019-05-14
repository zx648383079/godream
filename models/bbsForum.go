package models

type BbsForum struct {
	ID          uint
	Name        string
	Thumb       string
	Description string
	ParentID    int
	ThreadCount int
	PostCount   int
	Type        int
	Position    int
	CreatedAt   int
	UpdatedAt   int
}
