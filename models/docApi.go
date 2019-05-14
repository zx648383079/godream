package models

type DocApi struct {
	ID          uint
	Name        string
	Method      string
	Uri         string
	ProjectID   int
	Description string
	ParentID    int
	CreatedAt   int
	UpdatedAt   int
}
