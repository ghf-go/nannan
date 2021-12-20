package commonlogic

import (
	"context"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"strconv"
)

func GetGroupAll() map[int64]string {
	ret := map[int64]string{}
	m, e := logic.GetRedis().HGetAll(context.Background(), _redisGroupKey).Result()
	if e != nil {
		return ret
	}
	for k, v := range m {
		id, e := strconv.ParseInt(k, 10, 64)
		if e == nil {
			ret[id] = v
		}
	}
	return ret
}
func NewGroup(groupName string) {
	id := logic.GetTable(tb_system_group).InsertMap(db.Data{"group_name": groupName})
	if id > 0 {
		logic.GetRedis().HSet(context.Background(), _redisGroupKey, string(id), groupName)
	}
}