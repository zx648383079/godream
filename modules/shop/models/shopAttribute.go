package models

type ShopAttribute struct {
	ID           uint
	Name         string
	GroupID      int
	Type         int
	SearchType   int
	InputType    int
	DefaultValue string
	Position     int
}
