package models

type BlogTerm struct {
	ID          uint
	Name        string
	ParentID    int
	Keywords    string
	Description string
	Thumb       string
}
