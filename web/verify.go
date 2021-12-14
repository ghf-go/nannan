package web

import "regexp"

var (
	_regEmail =  regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)
	_regMobile =  regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")
	_regUrl =  regexp.MustCompile("/^(https?:\\/\\/)?([\\da-z\\.-]+)\\.([a-z\\.]{2,6})([\\/\\w \\.-]*)*\\/?$/")
	_regIpV4 =  regexp.MustCompile("/^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$/")
	_regIpV6 =  regexp.MustCompile("/^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/")
	_regDate = regexp.MustCompile("/^((19)|(20))\\d{2}\\-(0|1)\\d\\-[0-3]{1}\\d$/")
	_regTime = regexp.MustCompile("/^[0-2]{1}\\d:[0-5]{1}\\d:[0-5]{1}\\d$/")
	_regDateTime = regexp.MustCompile("/^((19)|(20))\\d{2}\\-(0|1)\\d\\-[0-3]{1}\\d [0-2]{1}\\d:[0-5]{1}\\d:[0-5]{1}\\d$/")
)

//是否是邮箱
func IsEmail(mail string) bool {
	return _regEmail.MatchString(mail)
}
func IsMobile(mobile string) bool {
	return _regMobile.MatchString(mobile)
}
func IsDate(date string) bool {
	return _regDate.MatchString(date)
}
func IsTime(time string) bool {
	return _regTime.MatchString(time)
}
func IsIPv4(ip string) bool  {
	return _regIpV4.MatchString(ip)
}
func IsIPv6(ip string) bool  {
	return _regIpV6.MatchString(ip)
}
func IsUrl(url string) bool  {
	return _regUrl.MatchString(url)
}
func IsDateTime(datetime string) bool {
	return _regDateTime.MatchString(datetime)
}
