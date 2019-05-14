package models

type ShopOrderActivity struct {
	ID        uint
	OrderID   int
	ProductID int
	Type      int
	Amount    float32
	Tag       string
	Name      string
}
