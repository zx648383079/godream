package models

type DocField struct {
	ID           uint
	Name         string
	Title        string
	IsRequired   int
	DefaultValue string
	Mock         string
	ParentID     int
	ApiID        int
	Kind         int
	Type         string
	Remark       string
	CreatedAt    int
	UpdatedAt    int
}
