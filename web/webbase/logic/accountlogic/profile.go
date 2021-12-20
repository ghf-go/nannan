package accountlogic

import (
	"context"
	"fmt"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"time"
)

const (
	redisKeyProfile    = "u:p:%d"
	redisExpireProfile = time.Second * 86400 * 365
	tb_user_profile    = "tb_user_profile"
)

func GetProfileByUid(uid int64) map[string]string {
	rd := logic.GetRedis()
	rk := getRedisKeyProfile(uid)
	ret, e := rd.HGetAll(context.Background(), rk).Result()
	if e == nil {
		return ret
	}
	var lis []modelUserProfile
	r := map[string]string{}
	if logic.CreateQuery(tb_user_profile).Where("user_id=?", uid).FetchAll(&lis) == nil {
		rdata := []interface{}{}
		for _, row := range lis {
			r[row.Key] = row.Val
			rdata = append(rdata, row.Key, row.Val)
		}
		rd.HSet(context.Background(), rk, rdata...)
	}
	return r
}

func SetProfile(uid int64, data map[string]interface{}) {
	for k, v := range data {
		if logic.CreateQuery(tb_user_profile).Where("user_id=? AND `key`=?", uid, k).Count("id") > 0 {
			logic.CreateQuery(tb_user_profile).Where("user_id=? AND `key`=?", uid, k).UpdateMap(db.Data{"val": v})
		} else {
			logic.GetTable(tb_user_profile).InsertMap(map[string]interface{}{
				"user_id": uid,
				"key":     k,
				"val":     v,
			})
		}
	}
	rd := logic.GetRedis()
	rk := getRedisKeyProfile(uid)
	var lis []modelUserProfile
	r := map[string]string{}
	if logic.CreateQuery(tb_user_profile).Where("user_id=?", uid).FetchAll(&lis) == nil {
		rdata := []interface{}{}
		for _, row := range lis {
			r[row.Key] = row.Val
			rdata = append(rdata, row.Key, row.Val)
		}
		rd.HSet(context.Background(), rk, rdata...)
	}
}

func getRedisKeyProfile(uid int64) string {
	return fmt.Sprintf(redisKeyProfile, uid)
}
