package commentlogic

import (
	"context"
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/webbase/logic"
	"strconv"
)

func Praise(uid int64, targetId int64, targetType int) bool {
	if IsPraise(uid, targetId, targetType) {
		return true
	}
	ret := logic.GetTable(tb_comment_praise).InsertMap(def.Data{
		"user_id":     uid,
		"target_id":   targetId,
		"target_type": targetType,
	}) > 0
	if ret {
		logic.GetRedis().HIncrBy(context.Background(), redisPraiseTotalKey(targetType), strconv.FormatInt(targetId, 10), 1)
		logic.GetRedis().HSet(context.Background(), redisPraiseKey(uid, targetType), strconv.FormatInt(targetId, 10), true)
	}
	return ret
}
func UnPraise(uid int64, targetId int64, targetType int) bool {
	if !IsPraise(uid, targetId, targetType) {
		return true
	}
	ret := logic.CreateQuery(tb_comment_praise).Where("user_id=? AND target_id=> AND target_type=?", uid, targetId, targetType).Delete() > 0
	if ret {
		logic.GetRedis().HDel(context.Background(), redisPraiseKey(uid, targetType), strconv.FormatInt(targetId, 10))
		logic.GetRedis().HIncrBy(context.Background(), redisPraiseTotalKey(targetType), strconv.FormatInt(targetId, 10), -1)
	}
	return ret
}
func IsPraise(uid int64, targetId int64, targetType int) bool {
	r, e := logic.GetRedis().HGet(context.Background(), redisPraiseKey(uid, targetType), strconv.FormatInt(targetId, 10)).Bool()
	if e != nil {
		return false
	}
	return r
}
func PraiseMap(uid int64, targetType int) map[int64]bool {
	ret := map[int64]bool{}
	r, e := logic.GetRedis().HGetAll(context.Background(), redisPraiseKey(uid, targetType)).Result()
	if e != nil {
		return ret
	}
	for k := range r {
		tid, e := strconv.ParseInt(k, 10, 64)
		if e == nil {
			ret[tid] = true
		}
	}
	return ret

}

func PraiseTotalMap(targetType int) map[int64]int64 {
	ret := map[int64]int64{}
	r, e := logic.GetRedis().HGetAll(context.Background(), redisPraiseTotalKey(targetType)).Result()
	if e != nil {
		return ret
	}
	for k, v := range r {
		tid, e := strconv.ParseInt(k, 10, 64)
		if e == nil {
			v, e := strconv.ParseInt(v, 10, 64)
			if e != nil {
				ret[tid] = 0
			} else {
				ret[tid] = v
			}

		}
	}
	return ret

}
