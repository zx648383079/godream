package models

type ShopOrderLog struct {
	ID        uint
	OrderID   int
	UserID    int
	Status    int
	Remark    string
	CreatedAt int
}
