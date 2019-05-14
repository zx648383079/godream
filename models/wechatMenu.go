package models

type WechatMenu struct {
	ID        uint
	Wid       int
	Name      string
	Type      string
	Content   string
	Pages     string
	ParentID  int
	CreatedAt int
	UpdatedAt int
}
