package models

type TplPageWeight struct {
	ID        uint
	PageID    int
	WeightID  int
	ParentID  int
	Position  int
	Title     string
	Content   string
	Settings  string
	IsShare   int
	CreatedAt int
	UpdatedAt int
}
