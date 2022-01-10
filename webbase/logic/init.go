package logic

import (
	"github.com/ghf-go/nannan/drivers/db_driver"
	"github.com/ghf-go/nannan/drivers/es_driver"
	"github.com/ghf-go/nannan/mod"
	"github.com/go-redis/redis/v8"
)

var (
	_redisConf = "default"
	_dbConf    = "default"
	_esConf    = ""
)

//是否使用es存储数据
func IsEsEnable() bool {
	return _esConf != ""
}
func RegisterEsName(esConf string) {
	_esConf = esConf
}
func RegisterDBName(dbConf string) {
	_dbConf = dbConf
}
func RegisterRedisName(redisConfName string) {
	_redisConf = redisConfName
}

func GetEsClient() *es_driver.EsClient {
	return mod.GetEsClient(_esConf)
}
func GetRedis() *redis.Client {
	return mod.GetRedis(_redisConf)
}
func GetDB() *db_driver.DBCon {
	return mod.NewDBClient(_dbConf)
}
func GetTable(table string) *db_driver.Table {
	return GetDB().Table(table)
}
func CreateQuery(table string) *db_driver.Query {
	return GetDB().Table(table).CreateQuery()
}
