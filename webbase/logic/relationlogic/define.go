package relationlogic

import "fmt"

const (
	FRIEND_STATUS_APPLY      = 10
	FRIEND_STATUS_WAIT_AUDIT = 20
	FRIEND_STATUS_OK         = 200
	tb_relation_follow       = "tb_relation_follow"
	tb_relation_friends      = "tb_relation_friends"
	tb_relation_blacklist    = "tb_relation_blacklist"

	redisKeyFollow      = "r:fo:%d"
	redisKeyBlackList   = "r:b:%s"
	redisKeyFriend      = "r:f:%s"
	redisFollowTotal    = "total"
	redisFollowFanTotal = "fans"
	redisBaclListTotal  = "total"
)

func getRedisFollowKey(uid int64) string {
	return fmt.Sprintf(redisKeyFollow, uid)
}
func getRedisBlackKey(uid int64) string {
	return fmt.Sprintf(redisKeyBlackList, uid)
}
func getRedisFriendKey(uid int64) string {
	return fmt.Sprintf(redisKeyFriend, uid)
}
