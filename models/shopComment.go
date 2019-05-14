package models

type ShopComment struct {
	ID        uint
	UserID    int
	ItemType  int
	ItemID    int
	Title     string
	Content   string
	Rank      int
	CreatedAt int
	UpdatedAt int
}
