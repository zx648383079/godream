package models

type WechatUser struct {
	ID            uint
	Openid        string
	Nickname      string
	Sex           int
	City          string
	Country       string
	Province      string
	Language      string
	Avatar        string
	SubscribeTime int
	UnionID       string
	Remark        string
	GroupID       int
	UpdatedAt     int
}
