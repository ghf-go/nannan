package relationlogic

import (
	"context"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
)

//身体添加好友
func ApplyFriend(uid, targetId int64, msg string) bool {
	switch StatFriend(uid, targetId) {
	case FRIEND_STATUS_OK, FRIEND_STATUS_APPLY:
		return true
	case FRIEND_STATUS_WAIT_AUDIT:
		return AuditFriend(uid, targetId, true)
	default:
		table := logic.GetTable(tb_relation_friends)
		if table.InsertMap(db.Data{
			"user_id":   uid,
			"target_id": targetId,
			"msg":       msg,
			"status":    FRIEND_STATUS_APPLY,
		}) > 0 && table.InsertMap(db.Data{
			"user_id":   targetId,
			"target_id": uid,
			"msg":       msg,
			"status":    FRIEND_STATUS_WAIT_AUDIT,
		}) > 0 {
			redis := logic.GetRedis()
			redis.HSet(context.Background(), getRedisFriendKey(uid), string(targetId), FRIEND_STATUS_APPLY)
			redis.HSet(context.Background(), getRedisFriendKey(targetId), string(uid), FRIEND_STATUS_WAIT_AUDIT)
			return true
		}
		table.CreateQuery().Where("(user_id=? AND target_id=?) OR (user_id=? AND target_id=?)", uid, targetId, targetId, uid).Delete()
		return false
	}
	return true
}

//审核好友申请
func AuditFriend(uid, targetId int64, isOk bool) bool {
	if StatFriend(uid, targetId) == FRIEND_STATUS_WAIT_AUDIT {
		redis := logic.GetRedis()
		table := logic.GetTable(tb_relation_friends)
		if isOk {
			if table.CreateQuery().Where("(user_id=? AND target_id=?) OR (user_id=? AND target_id=?)", uid, targetId, targetId, uid).UpdateMap(db.Data{"status": FRIEND_STATUS_OK}) > 0 {
				redis.HSet(context.Background(), getRedisFriendKey(uid), string(targetId), FRIEND_STATUS_OK)
				redis.HSet(context.Background(), getRedisFriendKey(targetId), string(uid), FRIEND_STATUS_OK)
				return true
			} else {
				return false
			}
		} else {
			table.CreateQuery().Where("(user_id=? AND target_id=?) OR (user_id=? AND target_id=?)", uid, targetId, targetId, uid).Delete()
			redis.HDel(context.Background(), getRedisFriendKey(uid), string(targetId))
			redis.HDel(context.Background(), getRedisFriendKey(targetId), string(uid))
		}
	}
	return true
}

//好友申请状态
func StatFriend(uid, targetId int64) int {
	r, e := logic.GetRedis().HGet(context.Background(), getRedisFriendKey(uid), string(targetId)).Int()
	if e != nil {
		return 0
	}
	return r
}

//是否是好友
func IsFriend(uid, targetId int64) bool {
	return StatFriend(uid, targetId) == FRIEND_STATUS_OK
}

//更新黑名单
func ReLoadFriendCacheByUid(uid int64) {

}

//更新全部的黑名单
func ReloadFriendCacheAll() {

}
