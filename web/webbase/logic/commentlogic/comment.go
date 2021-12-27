package commentlogic

import (
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/web/webbase/logic"
)

func NewComment(uid, targetid, parent_id int64, targetType int, content string) *CommentModel {
	id := logic.GetTable(tb_comment_reply).InsertMap(db.Data{
		"user_id":     uid,
		"target_id":   targetid,
		"parent_id":   parent_id,
		"target_type": targetType,
		"content":     content,
		"reply_count": 0,
	})
	if id > 0 {
		if parent_id > 0 {
			logic.CreateQuery(tb_comment_reply).Where("id=?", parent_id).UpdateMap(db.Data{"reply_count+": 1})
		}
		ret := &CommentModel{}
		if logic.CreateQuery(tb_comment_reply).Where("id=?", id).Frist(ret) == nil {
			return ret
		}
	}
	return nil
}
func CommentList(uid, targetid, parent_id int64, targetType, start, pageSize int) []*CommentModel {
	rows := []*CommentModel{}
	sql := "target_id=? AND parent_id=? AND target_type=?"
	args := []interface{}{targetid, parent_id, targetType}
	if uid > 0 {
		sql += " AND ((user_id=? AND `status`=?) OR `status`=?)"
		args = append(args, uid, STATUS_WAIT_AUDIT, STATUS_AUDIT)
	} else {
		sql += " AND `status`=?"
		args = append(args, STATUS_AUDIT)
	}
	logic.CreateQuery(tb_comment_reply).Where(sql, args...).Skip(start).Limit(pageSize).FetchAll(&rows)
	return rows
}
func CommentTotal(uid, targetid, parent_id int64, targetType int) int64 {
	sql := "target_id=? AND parent_id=? AND target_type=?"
	args := []interface{}{targetid, parent_id, targetType}
	if uid > 0 {
		sql += " AND ((user_id=? AND `status`=?) OR `status`=?)"
		args = append(args, uid, STATUS_WAIT_AUDIT, STATUS_AUDIT)
	} else {
		sql += " AND `status`=?"
		args = append(args, STATUS_AUDIT)
	}
	return logic.CreateQuery(tb_comment_reply).Where(sql, args...).Count("id")
}
