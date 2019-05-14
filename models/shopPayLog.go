package models

type ShopPayLog struct {
	ID        uint
	PaymentID int
	Type      int
	UserID    int
	Data      string
	Status    int
	Amount    float32
	CreatedAt int
	UpdatedAt int
}
