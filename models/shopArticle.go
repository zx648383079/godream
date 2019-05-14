package models

type ShopArticle struct {
	ID          uint
	CatID       int
	Title       string
	Keywords    string
	Thumb       string
	Description string
	Brief       string
	Url         string
	File        string
	Content     string
	CreatedAt   int
	UpdatedAt   int
}
