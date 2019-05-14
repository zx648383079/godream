package models

type ShopOrderGood struct {
	ID              uint
	OrderID         int
	GoodsID         int
	UserID          int
	Name            string
	SeriesNumber    string
	Thumb           string
	Amount          int
	Price           float32
	RefundID        int
	Status          int
	AfterSaleStatus int
	CommentID       int
}
