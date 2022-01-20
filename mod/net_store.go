package mod

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/drivers/store_driver"
	"github.com/ghf-go/nannan/gutils"
	"strings"
)

var (
	_net_store_map = map[string]def.NetStore{}
)

func GetNetStore(confName string) def.NetStore {
	if !strings.HasPrefix(confName, "store.") {
		confName = "store." + confName
	}
	if r, ok := _net_store_map[confName]; ok {
		return r
	}
	conf := GetConf(confName)
	switch conf.Scheme {
	case "ali":
	case "qiniu":
		q := &store_driver.Qiniu{
			AccessKey: conf.GetArgs("ak"),
			SecretKey: conf.GetArgs("sk"),
			Domain:    conf.Host,
			Bucket:    conf.Path,
		}
	default:
		Error("CDN 配置错误 %s", conf.Raw)
		gutils.Error(500, "CDN 配置错误")
	}
	return nil
}
