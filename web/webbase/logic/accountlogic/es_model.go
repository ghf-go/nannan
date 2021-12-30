package accountlogic

import "github.com/ghf-go/nannan/db"

type esFindUser struct {
	*db.EsResponseBase
	*db.EsResponseBaseDoc
	Data *UserInfo `json:"_source"`
}
type esMgetUser struct {
	*db.EsResponseBase
	Docs []struct {
		*db.EsResponseDocsBaseDoc
		Data *UserInfo `json:"_source"`
	} `json:"docs"`
}
