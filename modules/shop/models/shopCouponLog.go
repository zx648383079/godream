package models

type ShopCouponLog struct {
	ID           uint
	CouponID     int
	SerialNumber string
	UserID       int
	OrderID      int
	UsedAt       int
	CreatedAt    int
	UpdatedAt    int
}
