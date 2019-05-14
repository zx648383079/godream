package models

type ShopOrder struct {
	ID           uint
	SeriesNumber string
	UserID       int
	Status       int
	PaymentID    int
	PaymentName  string
	ShippingID   int
	InvoiceID    int
	ShippingName string
	GoodsAmount  float32
	OrderAmount  float32
	Discount     float32
	ShippingFee  float32
	PayFee       float32
	PayAt        int
	ShippingAt   int
	ReceiveAt    int
	FinishAt     int
	CreatedAt    int
	UpdatedAt    int
}
