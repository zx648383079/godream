package models

type Queries struct {
	Page     uint   `form:"page" json:"page"`
	PerPage  uint   `form:"per_page" json:"per_page"`
	Keywords string `form:"keywords" json:"keywords"`
}
