package models

type (
	Archives struct {
		Year  uint            `json:"year"`
		Items []*ArchivesItem `json:"items"`
	}
	ArchivesItem struct {
		Month string `json:"month"`
		Title string `json:"title"`
		ID    uint   `json:"id"`
	}
)
