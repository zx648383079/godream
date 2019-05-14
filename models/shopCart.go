package models

type ShopCart struct {
	ID               uint
	Type             int
	UserID           int
	GoodsID          int
	Amount           int
	Price            float32
	IsChecked        int
	SelectedActivity int
}
