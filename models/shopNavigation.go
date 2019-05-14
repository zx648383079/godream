package models

type ShopNavigation struct {
	ID       uint
	Type     string
	Name     string
	Url      string
	Target   string
	Visible  int
	Position int
}
