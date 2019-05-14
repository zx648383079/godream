package models

type BaseFeedback struct {
	ID        uint
	Name      string
	Email     string
	Phone     string
	Content   string
	Status    int
	CreatedAt int
	UpdatedAt int
}
