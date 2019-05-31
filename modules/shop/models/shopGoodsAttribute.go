package models

type ShopGoodsAttribute struct {
	ID          uint
	GoodsID     int
	AttributeID int
	Value       string
	Price       float32
}
