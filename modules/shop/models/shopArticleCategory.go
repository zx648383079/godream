package models

type ShopArticleCategory struct {
	ID          uint
	Name        string
	Keywords    string
	Description string
	ParentID    int
	Position    int
}
