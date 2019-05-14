package models

type WechatMedium struct {
	ID           uint
	Wid          int
	Type         string
	MaterialType int
	Title        string
	Thumb        string
	ShowCover    int
	OpenComment  int
	OnlyComment  int
	Content      string
	ParentID     int
	MediaID      string
	Url          string
	ExpiredAt    int
	CreatedAt    int
	UpdatedAt    int
}
