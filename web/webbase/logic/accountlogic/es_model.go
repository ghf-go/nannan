package accountlogic

import "github.com/ghf-go/nannan/db"

type EsUser struct {
	UserId      int64  `json:"user_id"`
	Role        int    `json:"role"`
	WxOpenID    string `json:"wx_open_id"`
	WxAvatar    string `json:"wx_avatar"`
	WxNickName  string `json:"wx_nick_name"`
	ProfileCity string `json:"profile_city"`
}

type esFindUser struct {
	*db.EsResponseBase
	*db.EsResponseBaseDoc
	Data *EsUser `json:"_source"`
}
type esMgetUser struct {
	*db.EsResponseBase
	Docs []struct {
		*db.EsResponseDocsBaseDoc
		Data *EsUser `json:"_source"`
	} `json:"docs"`
}
