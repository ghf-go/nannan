package relation

type reqFollow struct {
	TargetUid int64 `post:"uid" verify:"required"`
	IsDel     bool  `post:"is_del"`
}
type reqBlackList struct {
	TargetUid int64 `post:"uid" verify:"required"`
	IsDel     bool  `post:"is_del"`
}

type reqFriendApply struct {
	TargetUid int64  `post:"uid" verify:"required"`
	Msg       string `post:"msg" verify:"required"`
}
type reqFriendAudit struct {
	TargetUid int64 `post:"uid" verify:"required"`
	IsDel     bool  `post:"is_refuse"`
}
