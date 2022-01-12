package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/gutils"
)

var (
	_net_store_map = map[string]def.NetStore{}
)

func GetNetStore(confName string) def.NetStore {
	if r, ok := _net_store_map[confName]; ok {
		return r
	}
	conf := GetConf(confName)
	switch conf.Scheme {
	case "ali":
	case "qiniu":
	default:
		Error("CDN 配置错误 %s", conf.Raw)
		gutils.Error(500, "CDN 配置错误")
	}
	return nil
}
