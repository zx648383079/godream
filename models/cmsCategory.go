package models

type CmsCategory struct {
	ID               uint
	Name             string
	Title            string
	Type             int
	ModelID          int
	ParentID         int
	Keywords         string
	Description      string
	Image            string
	Content          string
	Url              string
	Position         int
	Groups           string
	CategoryTemplate string
	ListTemplate     string
	ShowTemplate     string
	Setting          string
	CreatedAt        int
	UpdatedAt        int
}
