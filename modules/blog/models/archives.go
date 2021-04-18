package models

type (
	Archives struct {
		Year  uint
		Items []*ArchivesItem
	}
	ArchivesItem struct {
		Month string
		Title string
		ID    uint
	}
)
