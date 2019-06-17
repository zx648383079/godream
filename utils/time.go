package utils

import "time"

// Now 获取当前时间的时间戳
func Now() int64 {
	return time.Now().Unix()
}

// StrToTime 字符串转化成时间戳
func StrToTime(format string, date string) int64 {
	t, _ := time.Parse(format, date)
	return t.Unix()
}

// Date 格式化时间
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
