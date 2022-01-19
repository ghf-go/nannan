package mod

import (
	"database/sql"
	"github.com/ghf-go/nannan/drivers/db_driver"
	"github.com/ghf-go/nannan/gutils"
	"strings"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var (
	_dbMap = map[string]*sql.DB{}
	//_dbConMap = map[string]*db_driver.DBCon{}
)

//新建数据库链接
func NewDB(confKeyName string) *sql.DB {
	if !strings.HasPrefix(confKeyName, "db.") {
		confKeyName = "db." + confKeyName
	}
	conf := GetConf(confKeyName)
	arrs := strings.Split(conf.Raw, "://")
	if len(arrs) != 2 {
		Error("数据库配置错误 %s -> %s", confKeyName, conf.Raw)
		gutils.Error(500, "数据库配置错误")
	}

	db, e := sql.Open(arrs[0], arrs[1])
	if e != nil {
		Error("数据库链接错误 (%s) %s", e.Error(), conf.Raw)
		gutils.Error(500, "数据库配置错误")
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Minute * 30)
	db.SetConnMaxLifetime(time.Hour * 2)
	return db
}

// 获取数据库链接
func GetDB(confKeyName string) *sql.DB {
	if r, ok := _dbMap[confKeyName]; ok {
		return r
	}
	db := NewDB(confKeyName)
	_dbMap[confKeyName] = db
	return db
}

// 获取数据库链接
func NewDBClient(confKeyName string) *db_driver.DBCon {
	return db_driver.NewCon(GetDB(confKeyName))
}

//
//// 获取链接
//func GetDBClient(confKeyName string) *db_driver.DBCon {
//	if r, ok := _dbConMap[confKeyName]; ok {
//		return r
//	}
//	db := NewDBClient(confKeyName)
//	_dbMap[confKeyName] = db
//	return db
//}
