package utils

import "time"

// 获取当前时间的时间戳
func Now() int64 {
	return time.Now().Unix()
}

func StrToTime(format string, date string) int64 {
	t, _ := time.Parse(format, date)
	return t.Unix()
}

func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
