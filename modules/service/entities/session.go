package entities

// Session Model
type Session struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	Remark      string `json:"remark"`
	UserId      int    `json:"user_id"`
	ServiceId   int    `json:"service_id"`
	Ip          string `json:"ip"`
	UserAgent   string `json:"user_agent"`
	Status      int32  `json:"status"`
	ServiceWord int    `json:"service_word"`
	UpdatedAt   int    `json:"updated_at"`
	CreatedAt   int    `json:"created_at"`
}

// TableName 表名
func (Session) TableName() string {
	return "service_session"
}

// SessionLog Model
type SessionLog struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	UserId    int    `json:"user_id"`
	SessionId int    `json:"session_id"`
	Remark    string `json:"remark"`
	Status    int32  `json:"status"`
	CreatedAt int    `json:"created_at"`
}

// TableName 表名
func (SessionLog) TableName() string {
	return "service_session_log"
}
