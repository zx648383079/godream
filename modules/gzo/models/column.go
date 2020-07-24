package models

// Column 表的列
type Column struct {
	Name          string
	Comment       string
	DataType      string
	Nullable      bool
	TitleCaseName string
	CamelCaseName string
	GoType        string
	Tag           string
}
