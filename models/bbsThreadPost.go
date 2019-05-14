package models

type BbsThreadPost struct {
	ID        uint
	Content   string
	ThreadID  int
	UserID    int
	Ip        string
	Grade     int
	CreatedAt int
	UpdatedAt int
}
