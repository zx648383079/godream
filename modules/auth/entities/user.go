package entities

import "time"

// User Model
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Sex       int32     `json:"sex"`
	Avatar    string    `json:"avatar"`
	Birthday  time.Time `json:"birthday"`
	Money     int       `json:"money"`
	Credits   int       `json:"credits"`
	ParentId  int       `json:"parent_id"`
	Token     string    `json:"token"`
	Status    int32     `json:"status"`
	UpdatedAt int       `json:"updated_at"`
	CreatedAt int       `json:"created_at"`
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
