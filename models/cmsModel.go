package models

type CmsModel struct {
	ID               uint
	Name             string
	Table            string
	Type             int
	Position         int
	CategoryTemplate string
	ListTemplate     string
	ShowTemplate     string
	Setting          string
}
