package models

type ShopOrderCoupon struct {
	ID       uint
	OrderID  int
	CouponID int
	Name     string
	Type     string
}
