package models

type ShopDelivery struct {
	ID              uint
	UserID          int
	OrderID         int
	Status          int
	ShippingID      int
	ShippingName    string
	ShippingFee     float32
	LogisticsNumber string
	Name            string
	RegionID        int
	RegionName      string
	Tel             string
	Address         string
	BestTime        string
	CreatedAt       int
	UpdatedAt       int
}
