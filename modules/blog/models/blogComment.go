package models

type BlogComment struct {
	ID        uint
	Content   string
	Name      string
	Email     string
	Url       string
	ParentID  int
	Position  int
	UserID    int
	BlogID    int
	Ip        string
	Agent     string
	Agree     int
	Disagree  int
	Approved  int
	CreatedAt int
}
