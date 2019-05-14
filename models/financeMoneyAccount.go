package models

type FinanceMoneyAccount struct {
	ID          uint
	Name        string
	Money       float32
	FrozenMoney float32
	Status      int
	Remark      string
	UserID      int
	DeletedAt   int
	CreatedAt   int
	UpdatedAt   int
}
