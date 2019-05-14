package models

type UserOauth struct {
	ID        uint
	UserID    int
	Nickname  string
	Vendor    string
	Identity  string
	Data      string
	CreatedAt int
}
