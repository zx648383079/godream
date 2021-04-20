package entities

// BlackWord 违禁词Model
type BlackWord struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	Words        string `json:"words"`
	ReplaceWords string `json:"replace_words"`
}

// TableName 表名
func (BlackWord) TableName() string {
	return "seo_black_word"
}
