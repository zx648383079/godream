package models

type CmsLinkageDatum struct {
	ID        uint
	LinkageID int
	Name      string
	ParentID  int
	Position  int
}
