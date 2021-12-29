package commonlogic

import "fmt"

const (
	tb_system_conf  = "tb_system_conf"
	tb_system_group = "tb_system_group"
	tb_system_tags  = "tb_system_tags"

	_esFeed = "feed"

	_redisConfKey  = "s:c:%d"
	_redisTagKey   = "s:t:%d"
	_redisGroupKey = "s:g:a"
)

func redisConfKey(groupid int64) string {
	return fmt.Sprintf(_redisConfKey, groupid)
}
func redisTagKey(groupid int64) string {
	return fmt.Sprintf(_redisTagKey, groupid)
}
