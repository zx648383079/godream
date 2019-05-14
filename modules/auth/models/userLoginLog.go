package models

type UserLoginLog struct {
	ID        uint
	Ip        string
	UserID    int
	User      string
	Status    int
	Mode      string
	CreatedAt int
}
