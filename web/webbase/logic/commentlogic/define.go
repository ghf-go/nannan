package commentlogic

import "fmt"

const (
	tb_comment_praise = "tp_comment_praise"
	tb_comment_score  = "tb_comment_score"
	tb_comment_reply  = "tb_comment_reply"

	_redisPraiseKey      = "c:p:%d:%d"
	_redisPraiseTotalKey = "c:p:t:%d:%d"

	_redisScoreKey    = "c:s:%d"
	PRAISE_TYPE_FEED  = 1
	PRAISE_TYPE_REPLY = 2
)

func redisPraiseKey(uid int64, targetType int) string {
	return fmt.Sprintf(_redisPraiseKey, uid, targetType)
}
func redisPraiseTotalKey(targetType int) string {
	return fmt.Sprintf(_redisPraiseTotalKey, 0, targetType)
}
