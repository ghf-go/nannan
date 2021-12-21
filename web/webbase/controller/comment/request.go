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

type reqFeedAdd struct {
	FeedTitle string  `post:"title"`
	FeedType  int     `post:"type" verify:"required"`
	FeedDesc  string  `post:"desc" verify:"required"`
	FeedImgs  string  `post:"imgs"`
	X         float64 `post:"x"`
	Y         float64 `post:"y"`
	City      string  `post:"city"`
	Ext       string  `post:"ext"`
	Content   string  `post:"content" verify:"required"`
}
