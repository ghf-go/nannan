package common

type ReqSendSmsCode struct {
	Mobile string `post:"mobile" verify:"required;mobile"`
	SendType int `post:"sendType" verify:"required"`
}
