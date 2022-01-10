package accountlogic

type modelUserEmail struct {
	ID     int64  `column:"id"`
	UserID int64  `column:"user_id"`
	Email  string `column:"email"`
}
type modelUserMobile struct {
	ID     int64  `column:"id"`
	UserID int64  `column:"user_id"`
	Mobile string `column:"mobile"`
}
type modelUserWx struct {
	ID     int64  `column:"id"`
	UserID int64  `column:"user_id"`
	OpenID string `column:"id"`
}
type modelUserProfile struct {
	ID     int64  `column:"id"`
	UserID int64  `column:"user_id"`
	Key    string `column:"key"`
	Val    string `column:"val"`
}
type modelUserPasswd struct {
	ID     int64  `column:"id"`
	Passwd string `column:"passwd"`
	Sign   string `column:"sign"`
}
