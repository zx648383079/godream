package models

type UserActionLog struct {
	ID        uint
	Ip        string
	UserID    int
	Action    string
	Remark    string
	CreatedAt int
}
