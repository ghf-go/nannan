package logic

import (
	"github.com/ghf-go/nannan/db"
	"github.com/ghf-go/nannan/drivers"
	"github.com/go-redis/redis/v8"
)

var (
	_redisConf = "default"
	_dbConf    = "default"
	_esConf    = ""
)

func RegisterEsName(esConf string) {
	_esConf = esConf
}
func RegisterDBName(dbConf string) {
	_dbConf = dbConf
}
func RegisterRedisName(redisConfName string) {
	_redisConf = redisConfName
}

func GetEsClient() *db.EsClient {
	return db.GetEsClient(_esConf)
}
func GetRedis() *redis.Client {
	return drivers.GetRedisByKey(_redisConf)
}
func GetDB() *db.DBCon {
	return db.GetDB(_dbConf)
}
func GetTable(table string) *db.Table {
	return GetDB().Table(table)
}
func CreateQuery(table string) *db.Query {
	return GetDB().Table(table).CreateQuery()
}
