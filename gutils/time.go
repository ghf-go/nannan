package gutils

import "time"

//日期格式化
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func NewAddDay(day int) time.Time {
	return time.Now().Add(time.Hour * 24 * time.Duration(day))
}
