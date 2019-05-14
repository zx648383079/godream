package models

type VisitLog struct {
	ID             uint
	Ip             string
	Browser        string
	BrowserVersion string
	Os             string
	OsVersion      string
	Url            string
	Referer        string
	Agent          string
	Country        string
	Region         string
	City           string
	UserID         int
	UserName       string
	Latitude       string
	Longitude      string
	CreatedAt      int
}
