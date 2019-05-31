package models

type ShopOrderRefund struct {
	ID           uint
	UserID       int
	OrderID      int
	OrderGoodsID int
	GoodsID      int
	ProductID    int
	Title        string
	Amount       int
	Type         int
	Status       int
	Reason       string
	Description  string
	Evidence     string
	Explanation  string
	Money        float32
	OrderPrice   float32
	Freight      int
	CreatedAt    int
	UpdatedAt    int
}
