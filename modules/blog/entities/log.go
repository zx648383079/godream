package entities

// BlogLog Model
type BlogLog struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Type      int32
	IdValue   int
	UserId    int
	Action    int
	CreatedAt int
}

// TableName 表名
func (BlogLog) TableName() string {
	return "blog_log"
}
