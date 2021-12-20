package commentlogic

import "fmt"

const (
	tb_comment_praise = "tp_comment_praise"
	tb_comment_score  = "tb_comment_score"
	tb_comment_reply  = "tb_comment_reply"
	tb_common_upload  = "tb_common_upload"

	_redisPraiseKey      = "c:p:%d:%d"
	_redisPraiseTotalKey = "c:p:t:%d:%d"
	_redisUploadKey      = "c:u:%s"

	_redisScoreKey    = "c:s:%d"
	STATUS_WAIT_AUDIT = 0
	STATUS_AUDIT      = 100
	PRAISE_TYPE_FEED  = 1
	PRAISE_TYPE_REPLY = 2
)

func redisPraiseKey(uid int64, targetType int) string {
	return fmt.Sprintf(_redisPraiseKey, uid, targetType)
}
func redisPraiseTotalKey(targetType int) string {
	return fmt.Sprintf(_redisPraiseTotalKey, 0, targetType)
}

func redisUploadKey(key string) string {
	return fmt.Sprintf(_redisUploadKey, key)
}
