package models

type CmsModelField struct {
	ID           uint
	Name         string
	Field        string
	ModelID      int
	Type         string
	Length       int
	Position     int
	FormType     int
	IsMain       int
	IsRequired   int
	IsDisable    int
	IsSystem     int
	Match        string
	TipMessage   string
	ErrorMessage string
	Setting      string
}
