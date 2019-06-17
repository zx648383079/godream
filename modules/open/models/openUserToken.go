package models

// OpenUserToken 开放账户下的用户登录信息
type OpenUserToken struct {
	ID         uint
	UserID     int
	PlatformID int
	Token      string
	ExpiredAt  int
	CreatedAt  int
	UpdatedAt  int
}
