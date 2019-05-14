package models

type DocPage struct {
	ID        uint
	Name      string
	ProjectID int
	ParentID  int
	Content   string
	CreatedAt int
	UpdatedAt int
}
