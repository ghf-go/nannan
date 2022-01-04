package gutils

// 设置字符串最大长度
func StrMaxLen(data string, maxline int) string {
	sub := []rune(data)
	if len(sub) > maxline {
		return string(sub[:maxline])
	}
	return data
}
