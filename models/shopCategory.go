package models

type ShopCategory struct {
	ID          uint
	Name        string
	Keywords    string
	Description string
	Icon        string
	Banner      string
	AppBanner   string
	ParentID    int
	Position    int
}
