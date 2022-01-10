package relationlogic

import (
	"context"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

//关注
func Follow(uid, targetId int64) bool {
	if IsFollow(uid, targetId) {
		return true
	}
	isOk := logic.GetTable(tb_relation_follow).InsertMap(map[string]interface{}{
		"user_id":   uid,
		"target_id": targetId,
	}) > 0
	if isOk {
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(uid), redisFollowTotal, 1)
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(targetId), redisFollowFanTotal, 1)
		logic.GetRedis().HSet(context.Background(), getRedisFollowKey(uid), strconv.FormatInt(targetId, 10), true)
	}
	return isOk
}

//取消关注
func UnFollow(uid, targetId int64) bool {
	if !IsFollow(uid, targetId) {
		return true
	}
	isOk := logic.CreateQuery(tb_relation_follow).Where("user_id=? AND target_id=?", uid, targetId).Delete() > 0
	if isOk {
		logic.GetRedis().HDel(context.Background(), getRedisFollowKey(uid), strconv.FormatInt(targetId, 10))
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(uid), redisFollowTotal, -1)
		logic.GetRedis().HIncrBy(context.Background(), getRedisFollowKey(targetId), redisFollowFanTotal, -1)
	}
	return isOk
}
func GetFanCount(uid int64) int64 {
	r, e := logic.GetRedis().HGet(context.Background(), getRedisFollowKey(uid), redisFollowFanTotal).Int64()
	if e != nil {
		return 0
	}
	return r
}
func GetFollowCount(uid int64) int64 {
	r, e := logic.GetRedis().HGet(context.Background(), getRedisFollowKey(uid), redisFollowTotal).Int64()
	if e != nil {
		return 0
	}
	return r
}

//获取关注列表
func GetFollowIDList(uid int64) []int64 {
	ret := []int64{}
	l, e := logic.GetRedis().HGetAll(context.Background(), getRedisFollowKey(uid)).Result()
	if e == nil {
		for k := range l {
			if k == redisFollowTotal || k == redisFollowFanTotal {
				continue
			}
			targetId, e := strconv.ParseInt(k, 10, 64)
			if e == nil {
				ret = append(ret, targetId)
			}
		}
	}
	return ret
}

//是否关注
func IsFollow(uid, targetId int64) bool {
	rk := getRedisFollowKey(uid)
	redis := logic.GetRedis()
	isFolow, e := redis.HGet(context.Background(), rk, strconv.FormatInt(targetId, 10)).Bool()
	if e != nil {
		return false
	}

	return isFolow
	//isFolow = logic.CreateQuery(tb_relation_follow).Where("user_id=? AND target_id=?", uid, targetId).Count("id") > 0
	//
	//redis.HSet(context.Background(), rk, string(targetId), isFolow)
	//return isFolow
}

//更新黑名单
func ReLoadFollowCacheByUid(uid int64) {

}

//更新全部的黑名单
func ReloadFollowCacheAll() {

}
