package comment

type reqNewComment struct {
	TargetID   int64  `post:"target_id" verify:"required"`
	TargetType int    `post:"target_type" verify:"required"`
	ParentID   int64  `post:"parent_id"`
	Content    string `post:"content" verify:"required"`
}

type reqCommentList struct {
	TargetID   int64 `post:"target_id" verify:"required"`
	TargetType int   `post:"target_type" verify:"required"`
	ParentID   int64 `post:"parent_id"`
	Start      int   `post:"start"`
	PageSize   int   `post:"page_size"`
}
type reqPraise struct {
	TargetID   int64 `post:"target_id" verify:"required"`
	TargetType int   `post:"target_type" verify:"required"`
	IsUnPraise bool  `post:"is_del"`
}
