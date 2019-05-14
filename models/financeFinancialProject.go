package models

import "time"

type FinanceFinancialProject struct {
	ID             uint
	Name           string
	Alias          string
	Money          float32
	AccountID      int
	Earnings       float32
	StartAt        time.Time
	EndAt          time.Time
	EarningsNumber float32
	ProductID      int
	Status         int
	DeletedAt      int
	Color          int
	Remark         string
	UserID         int
	CreatedAt      int
	UpdatedAt      int
}
