package models

import "time"

type FinanceLog struct {
	ID          uint
	Type        int
	Money       float32
	FrozenMoney float32
	AccountID   int
	ChannelID   int
	ProjectID   int
	BudgetID    int
	Remark      string
	HappenedAt  time.Time
	UserID      int
	CreatedAt   int
	UpdatedAt   int
}
