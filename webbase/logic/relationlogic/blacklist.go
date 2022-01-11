package relationlogic

import (
	"context"
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

//添加到黑名单
func AddBlackList(uid, targetid int64) bool {
	if InBlackList(uid, targetid) {
		return true
	}
	isOk := logic.GetTable(tb_relation_blacklist).InsertMap(def.Data{
		"user_id":        uid,
		"target_user_id": targetid,
	}) > 0
	if isOk {
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(uid), redisBaclListTotal, 1)
		logic.GetRedis().HSet(context.Background(), getRedisBlackKey(uid), strconv.FormatInt(targetid, 10), isOk)
	}
	return isOk
}

//删除黑名单
func DelBlackList(uid, targetid int64) bool {
	if !InBlackList(uid, targetid) {
		return true
	}
	isOk := logic.CreateQuery(tb_relation_blacklist).Where("user_id=? AND target_user_id=?", uid, targetid).Delete() > 0
	if isOk {
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(uid), redisBaclListTotal, -1)
		logic.GetRedis().HDel(context.Background(), getRedisBlackKey(uid), strconv.FormatInt(targetid, 10))
	}
	return isOk
}

//是否在黑名单中
func InBlackList(uid, targetid int64) bool {
	r, e := logic.GetRedis().HGet(context.Background(), getRedisBlackKey(uid), strconv.FormatInt(targetid, 10)).Bool()
	if e != nil {
		return false
	}
	return r
}
