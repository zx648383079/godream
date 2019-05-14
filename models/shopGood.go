package models

type ShopGood struct {
	ID               uint
	CatID            int
	BrandID          int
	Name             string
	SeriesNumber     string
	Keywords         string
	Thumb            string
	Picture          string
	Description      string
	Brief            string
	Content          string
	Price            float32
	MarketPrice      float32
	Stock            int
	AttributeGroupID int
	Weight           float32
	ShippingID       int
	Sales            int
	IsBest           int
	IsHot            int
	IsNew            int
	Status           int
	DeletedAt        int
	CreatedAt        int
	UpdatedAt        int
}
