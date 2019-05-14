package models

type ShopAd struct {
	ID         uint
	Name       string
	PositionID int
	Type       int
	Url        string
	Content    string
	StartAt    int
	EndAt      int
	CreatedAt  int
	UpdatedAt  int
}
