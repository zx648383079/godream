package models

// OpenPlatform 开放账户管理
type OpenPlatform struct {
	ID          uint
	UserID      int
	Name        string
	Type        int
	Domain      string
	Appid       string
	Secret      string
	SignType    int
	SignKey     string
	EncryptType int
	PublicKey   string
	Status      int
	CreatedAt   int
	UpdatedAt   int
}
