package models

type ShopCoupon struct {
	ID          uint
	Name        string
	Thumb       string
	Type        int
	Rule        int
	RuleValue   int
	MinMoney    float32
	Money       float32
	SendType    int
	SendValue   int
	EveryAmount int
	StartAt     int
	EndAt       int
	CreatedAt   int
	UpdatedAt   int
}
