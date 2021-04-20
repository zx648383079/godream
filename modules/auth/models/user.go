package models

type UserSimple struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Sex    int32  `json:"sex"`
	Avatar string `json:"avatar"`
}

// TableName 表名
func (UserSimple) TableName() string {
	return "user"
}
