package models

type UserLoginQr struct {
	ID        uint
	UserID    int
	Token     string
	Status    int
	ExpiredAt int
	CreatedAt int
	UpdatedAt int
}
