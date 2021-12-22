package common

type ReqSendSmsCode struct {
	Mobile   string `post:"mobile" verify:"required;mobile"`
	SendType int    `post:"sendType" verify:"required"`
}

type reqBaseGroup struct {
	GroupID int64 `post:"group_id" verify:"required"`
}

type reqGroupNew struct {
	GroupName string `post:"group" verify:"required"`
}
type reqNewTag struct {
	TagName string `post:"tag" verify:"required"`
	GroupID int64  `post:"group_id" verify:"required"`
}
type reqNewConf struct {
	Key     string `post:"key" verify:"required"`
	GroupID int64  `post:"group_id" verify:"required"`
	Desc    string `post:"desc" verify:"required"`
	Val     string `post:"val" verify:"required"`
	ValType int    `post:"val_type" verify:"required"`
}
