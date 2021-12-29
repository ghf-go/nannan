package gutils

import "strconv"

// 转字符串
func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}

//转字符串
func SlicInt64String(src []int64) []string {
	ret := []string{}
	for _, v := range src {
		ret = append(ret, Int64String(v))
	}
	return ret
}
