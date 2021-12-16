package drivers

import (
	"database/sql"
	"github.com/ghf-go/nannan/gconf"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

var (
	_dbMap = map[string]*sql.DB{}
)

func GetDBCon(key string) *sql.DB {
	if con ,ok := _dbMap[key];ok{
		return con
	}
	conf := gconf.GetRawConf(key)
	arrs := strings.Split(conf,"://")
	if len(arrs) !=2{
		panic("数据库配置错误")
	}
	db,e := sql.Open(arrs[0],arrs[1])
	if e != nil{
		panic(e)
	}
	_dbMap[key] = db
	return db
}
