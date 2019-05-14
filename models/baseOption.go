package models

type BaseOption struct {
	ID           uint
	Name         string
	Code         string
	ParentID     int
	Type         string
	Visibility   int
	DefaultValue string
	Value        string
	Position     int
}
