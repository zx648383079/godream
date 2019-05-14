package models

type OpenUserToken struct {
	ID         uint
	UserID     int
	PlatformID int
	Token      string
	ExpiredAt  int
	CreatedAt  int
	UpdatedAt  int
}
