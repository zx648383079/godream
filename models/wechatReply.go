package models

type WechatReply struct {
	ID        uint
	Wid       int
	Event     string
	Keywords  string
	Match     int
	Content   string
	Type      string
	CreatedAt int
	UpdatedAt int
}
