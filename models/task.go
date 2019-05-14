package models

type Task struct {
	ID          uint
	UserID      int
	ParentID    int
	Name        string
	Description string
	Status      int
	EveryTime   int
	TimeLength  int
	CreatedAt   int
	UpdatedAt   int
}
