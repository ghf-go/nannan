package gutils

import (
	"github.com/ghf-go/nannan/drivers/netstore"
	"github.com/ghf-go/nannan/gconf"
)

var (
	_storeMap = map[string]netstore.NetStore{}
)

func GetNetStore(name string) netstore.NetStore {
	if r, ok := _storeMap[name]; ok {
		return r
	}
	conf := gconf.GetConf("store." + name)
	switch conf.GetScheme() {
	case "ali":
		f := netstore.AliOss{
			Endpoint:        "",
			AccessKeyId:     "",
			AccessKeySecret: "",
			BucketName:      "",
			CdnDomain:       "",
			CdnType:         "",
			CdnSecret:       "",
			CdnExpire:       0,
		}
		_storeMap[name] = f
		return f
	}
	Error(123, "配置错误")
	return nil
}
