package models

type WechatMessageHistory struct {
	ID        uint
	Wid       int
	Rid       int
	Kid       int
	From      string
	To        string
	Message   string
	Type      string
	Mark      int
	CreatedAt int
}
