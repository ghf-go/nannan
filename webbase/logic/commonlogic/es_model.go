package commonlogic

type EsFeed struct {
	FeedId      int64  `json:"feed_id"`
	UserId      int64  `json:"user_id"`
	Role        int    `json:"role"`
	WxOpenID    string `json:"wx_open_id"`
	WxAvatar    string `json:"wx_avatar"`
	WxNickName  string `json:"wx_nick_name"`
	ProfileCity string `json:"profile_city"`
}
