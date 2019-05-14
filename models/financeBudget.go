package models

type FinanceBudget struct {
	ID        uint
	Name      string
	Budget    float32
	Spent     float32
	Cycle     int
	UserID    int
	DeletedAt int
	CreatedAt int
	UpdatedAt int
}
