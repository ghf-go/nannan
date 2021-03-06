package account

type reqLogin struct {
	Name   string `post:"login_name" verify:"required"`
	Passws string `post:"passwd"`
	Code   string `post:"code"`
}
type reqLogH5Wx struct {
	Code string `get:"code" verify:"required"`
}
type reqSetPass struct {
	Pass string `post:"pass" verify:"required"`
}

type reqBindMobile struct {
	Mobile string `post:"mobile" verify:"required;mobile"`
	Code   string `post:"code" verigy:"required"`
}
type reqBaseUid struct {
	Uid int64 `post:"uid"`
}

type reqUserProfile struct {
	Gender   string `post:"gender"`
	Province string `post:"province"`
	City     string `post:"city"`
	Birthday string `post:"birthday"`
	Avatar   string `post:"avatar"`
	NickName string `post:"nick_name"`
}
