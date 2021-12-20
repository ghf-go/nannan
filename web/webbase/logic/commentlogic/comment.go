package commentlogic

import (
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
)

func NewComment(uid, targetid, parent_id int64, targetType int, content string) {
	if logic.GetTable(tb_comment_reply).InsertMap(db.Data{
		"user_id":     uid,
		"target_id":   targetid,
		"parent_id":   parent_id,
		"target_type": targetType,
		"content":     content,
		"reply_count": 0,
	}) > 0 {
		if parent_id > 0 {
			logic.CreateQuery(tb_comment_reply).Where("id=?", parent_id).UpdateMap(db.Data{"reply_count+": 1})
		}
	}
}
func CommentList(targetid, parent_id int64, targetType, start, pageSize int) []CommentModel {
	rows := []CommentModel{}
	logic.CreateQuery(tb_comment_reply).Where("target_id=? AND parent_id=? AND target_type=?", targetid, parent_id, targetType).Skip(start).Limit(pageSize).FetchAll(&rows)
	return rows
}
func CommentTotal(targetid, parent_id int64, targetType int) int64 {
	return logic.CreateQuery(tb_comment_reply).Where("target_id=? AND parent_id=? AND target_type=?", targetid, parent_id, targetType).Count("id")
}
