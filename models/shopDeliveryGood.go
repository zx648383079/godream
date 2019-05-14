package models

type ShopDeliveryGood struct {
	ID           uint
	DeliveryID   int
	OrderGoodsID int
	GoodsID      int
	Name         string
	SeriesNumber string
	Amount       int
}
