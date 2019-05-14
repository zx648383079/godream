package models

type CheckIn struct {
	ID        uint
	UserID    int
	Type      int
	Running   int
	Money     int
	Ip        string
	Method    int
	CreatedAt int
}
