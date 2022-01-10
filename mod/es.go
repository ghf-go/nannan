package mod

import (
	"github.com/ghf-go/nannan/drivers/es_driver"
	"strings"
)

var (
	_esMap = map[string]*es_driver.EsClient{}
)

// 创建Es客户端
func NewEsClient(confKeyName string) *es_driver.EsClient {
	conf := GetConf(confKeyName)
	hosts := strings.Split(conf.Path, ",")
	r := &es_driver.EsClient{
		Hosts:     hosts,
		DbName:    conf.Host,
		ReqIndex:  0,
		HostCount: len(hosts),
	}
	return r
}

// 获取Es 客户端
func GetEsClient(confKeyName string) *es_driver.EsClient {
	if r, ok := _esMap[confKeyName]; ok {
		return r
	}
	db := NewEsClient(confKeyName)
	_esMap[confKeyName] = db
	return db
}
