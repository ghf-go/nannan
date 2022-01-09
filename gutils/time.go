package gutils

import "time"

//日期格式化
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func NewAddDay(day int) time.Time {
	return time.Now().Add(time.Hour * 24 * time.Duration(day))
}

//当前时间
func FormatCurTimeMicroSeconds() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}

//获取当前日期的时间戳
func CurDateUnix() int64 {
	t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return t.Unix()
}
func CurDate() string {
	return time.Now().Format("2006-01-02")
}
