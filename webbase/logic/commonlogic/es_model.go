package commonlogic

import (
	"github.com/ghf-go/nannan/drivers/es_driver"
)

type EsFeed struct {
	FeedId      int64  `json:"feed_id"`
	UserId      int64  `json:"user_id"`
	Role        int    `json:"role"`
	WxOpenID    string `json:"wx_open_id"`
	WxAvatar    string `json:"wx_avatar"`
	WxNickName  string `json:"wx_nick_name"`
	ProfileCity string `json:"profile_city"`
}

type esFindFeed struct {
	*es_driver.EsResponseBase
	*es_driver.EsResponseBaseDoc
	Data *EsFeed `json:"_source"`
}
type esMgetFeed struct {
	*es_driver.EsResponseBase
	Docs []struct {
		*es_driver.EsResponseDocsBaseDoc
		Data *EsFeed `json:"_source"`
	} `json:"docs"`
}
