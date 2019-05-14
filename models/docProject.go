package models

type DocProject struct {
	ID          uint
	UserID      int
	Name        string
	Type        int
	Description string
	Environment string
	Status      int
	DeletedAt   int
	CreatedAt   int
	UpdatedAt   int
}
