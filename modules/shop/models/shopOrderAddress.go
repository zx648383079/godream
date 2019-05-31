package models

type ShopOrderAddress struct {
	ID         uint
	OrderID    int
	Name       string
	RegionID   int
	RegionName string
	Tel        string
	Address    string
	BestTime   string
}
