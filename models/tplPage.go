package models

type TplPage struct {
	ID          uint
	SiteID      int
	Type        int
	Name        string
	Title       string
	Keywords    string
	Thumb       string
	Description string
	Template    string
	Settings    string
	Position    int
	DeletedAt   int
	CreatedAt   int
	UpdatedAt   int
}
