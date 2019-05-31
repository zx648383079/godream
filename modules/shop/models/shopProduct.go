package models

type ShopProduct struct {
	ID           uint
	GoodsID      int
	Price        float32
	MarketPrice  float32
	Stock        int
	Weight       float32
	SeriesNumber string
	Attributes   string
}
