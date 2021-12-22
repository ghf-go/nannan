package commonlogic

import (
	"context"
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
	"strconv"
)

func GetAllTagByGroupID(groupid int64) map[int64]string {
	ret := map[int64]string{}
	d, e := logic.GetRedis().HGetAll(context.Background(), redisTagKey(groupid)).Result()
	if e != nil {
		return ret
	}
	for k, v := range d {
		id, e := strconv.ParseInt(k, 10, 64)
		if e == nil {
			ret[id] = v
		}
	}
	return ret
}
func NewTag(groupid int64, tagname string) {
	id := logic.GetTable(tb_system_tags).InsertMap(db.Data{
		"group_id": groupid,
		"tag_name": tagname,
	})
	if id > 0 {
		logic.GetRedis().HSet(context.Background(), redisTagKey(groupid), strconv.FormatInt(id, 10), tagname)
	}
}
