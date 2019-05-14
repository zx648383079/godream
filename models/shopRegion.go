package models

type ShopRegion struct {
	ID       uint
	Name     string
	ParentID int
	FullName string
}
