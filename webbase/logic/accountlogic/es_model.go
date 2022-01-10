package accountlogic

import (
	"github.com/ghf-go/nannan/drivers/es_driver"
)

type esFindUser struct {
	*es_driver.EsResponseBase
	*es_driver.EsResponseBaseDoc
	Data *UserInfo `json:"_source"`
}
type esMgetUser struct {
	*es_driver.EsResponseBase
	Docs []struct {
		*es_driver.EsResponseDocsBaseDoc
		Data *UserInfo `json:"_source"`
	} `json:"docs"`
}
