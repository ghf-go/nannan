package commentlogic

import "time"

type CommentModel struct {
	ID         int64     `column:"id" json:"id"`
	UserId     int64     `column:"user_id" json:"user_id"`
	TargetId   int64     `column:"target_id" json:"target_id"`
	ParentId   int64     `column:"parent_id" json:"parent_id"`
	TargetType int       `column:"target_type" json:"target_type"`
	Content    string    `column:"content" json:"content"`
	ReplyCount int       `column:"reply_count" json:"reply_count"`
	CreateAt   time.Time `column:"create_at" json:"create_at"`
}
