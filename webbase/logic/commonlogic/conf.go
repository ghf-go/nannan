package commonlogic

import (
	"context"
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/webbase/logic"
)

func GetConfByGroupID(groupid int64) map[string]string {
	r, e := logic.GetRedis().HGetAll(context.Background(), redisConfKey(groupid)).Result()
	if e != nil {
		return map[string]string{}
	}
	return r
}
func NewConf(group_id int64, val_type int, key, desc, val string) {
	id := logic.GetTable(tb_system_conf).InsertMap(def.Data{
		"group_id": group_id,
		"key":      key,
		"desc":     desc,
		"val":      val,
		"val_type": val_type,
	})
	if id > 0 {
		logic.GetRedis().HSet(context.Background(), redisConfKey(group_id), key, val)
	}
}
