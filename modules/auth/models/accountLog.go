package models

type AccountLog struct {
	ID        uint
	UserID    int
	Type      int
	ItemID    int
	Money     int
	Status    int
	Remark    string
	CreatedAt int
	UpdatedAt int
}
