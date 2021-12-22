package drivers

import (
	"database/sql"
	"github.com/ghf-go/nannan/gconf"
	"github.com/ghf-go/nannan/glog"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var (
	_dbMap = map[string]*sql.DB{}
)

func GetDBCon(key string) *sql.DB {
	if con, ok := _dbMap[key]; ok {
		return con
	}
	conf := gconf.GetRawConf("db." + key)
	arrs := strings.Split(conf, "://")
	if len(arrs) != 2 {
		glog.Error("数据库配置错误")
		panic("数据库配置错误")
	}
	db, e := sql.Open(arrs[0], arrs[1])
	if e != nil {
		glog.Error("数据库链接错误%s", e.Error())
		panic(e)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Minute * 30)
	db.SetConnMaxLifetime(time.Hour * 2)

	_dbMap[key] = db
	return db
}
