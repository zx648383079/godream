package models

type ShopShippingGroup struct {
	ID            uint
	ShippingID    int
	FirstStep     float32
	FirstFee      float32
	Additional    float32
	AdditionalFee float32
	FreeStep      float32
}
