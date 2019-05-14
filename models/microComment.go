package models

type MicroComment struct {
	ID        uint
	Content   string
	ParentID  int
	UserID    int
	MicroID   int
	Agree     int
	Disagree  int
	CreatedAt int
}
