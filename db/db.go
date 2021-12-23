package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ghf-go/nannan/drivers"
	"github.com/ghf-go/nannan/glog"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Data map[string]interface{}
type DBCon struct {
	db *sql.DB
	tx *sql.Tx
}

func GetDB(conName string) *DBCon {
	return &DBCon{
		db: drivers.GetDBCon(conName),
	}
}
func (dbc *DBCon) Close() {
	dbc.Commit()
	dbc = nil
}
func (dbc *DBCon) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	glog.Debug("sql -> %s %v", sql, args)
	if dbc.tx == nil {
		return dbc.db.Query(sql, args...)
	}
	return dbc.tx.Query(sql, args...)
}

func (dbc *DBCon) Exec(sql string, args ...interface{}) (sql.Result, error) {
	glog.Debug("sql -> %s %v", sql, args)
	if dbc.tx == nil {
		return dbc.db.Exec(sql, args...)
	}
	return dbc.tx.Exec(sql, args...)
}
func (dbc *DBCon) FindSql(obj interface{}, sql string, args ...interface{}) error {
	rows, e := dbc.Query(sql, args)
	if e != nil {
		return e
	}
	return saveObj(rows, obj)
}
func (dbc *DBCon) FetchSql(obj interface{}, sql string, args ...interface{}) error {
	rows, e := dbc.Query(sql, args)
	if e != nil {
		return e
	}
	return saveObjList(rows, obj)
}
func (dbc *DBCon) DeleteSql(sql string, args ...interface{}) int64 {
	r, e := dbc.Exec(sql, args...)
	if e != nil {
		panic(e)
	}
	n, e := r.RowsAffected()
	if e != nil {
		panic(e)
	}
	return n
}
func (dbc *DBCon) InsertByMap(tableName string, data Data) int64 {
	keys := []string{}
	args := []interface{}{}
	vs := []string{}
	for k, v := range data {
		keys = append(keys, fmt.Sprintf("`%s`", k))
		args = append(args, v)
		vs = append(vs, "?")
	}
	r, e := dbc.Exec(fmt.Sprintf("INSERT INTO `%s`(%s) VALUES(%s)", tableName, strings.Join(keys, ","), strings.Join(vs, ",")), args...)
	if e != nil {
		panic(e)
	}
	id, e := r.LastInsertId()
	if e != nil {
		panic(e)
	}
	return id
}
func (dbc *DBCon) InsertBatchSql(sql string, args ...interface{}) int64 {
	r, e := dbc.Exec(sql, args...)
	if e != nil {
		panic(e)
	}
	n, e := r.RowsAffected()
	if e != nil {
		panic(e)
	}
	return n
}
func (dbc *DBCon) InsertSql(sql string, args ...interface{}) int64 {
	r, e := dbc.Exec(sql, args...)
	if e != nil {
		panic(e)
	}
	n, e := r.LastInsertId()
	if e != nil {
		panic(e)
	}
	return n
}
func (dbc *DBCon) UpdateSql(sql string, args ...interface{}) int64 {
	r, e := dbc.Exec(sql, args...)
	if e != nil {
		panic(e)
	}
	n, e := r.RowsAffected()
	if e != nil {
		panic(e)
	}
	return n
}
func (dbc *DBCon) QueryNum(sql string, args ...interface{}) int64 {
	r, e := dbc.Query(sql, args...)
	if e != nil {
		panic(e)
	}
	if r.Next() {
		var ret int64
		e = r.Scan(&ret)
		if e != nil {
			panic(e)
		}
		return ret
	}
	return 0
}
func (dbc *DBCon) Table(tbname string) *Table {
	return &Table{dbc, tbname}
}
func (dbc *DBCon) Begin() {
	if dbc.tx != nil {
		return
	}
	tc, e := dbc.db.Begin()
	if e != nil {
		panic(e)
	}
	dbc.tx = tc
}
func (dbc *DBCon) Commit() {
	if dbc.tx == nil {
		panic(errors.New("不在事务中"))
	}
	if dbc.tx.Commit() != nil {
		panic(errors.New("提交事务失败"))
	}
	dbc.tx = nil
}
func (dbc *DBCon) Rollback() {
	if dbc.tx == nil {
		panic(errors.New("不在事务中"))
	}
	if dbc.tx.Rollback() != nil {
		panic(errors.New("提交事务失败"))
	}
	dbc.tx = nil
}
func (dbc *DBCon) Tx(call func(con *DBCon)) {
	dbc.Begin()
	defer func() {
		if e := recover(); e != nil {
			dbc.Rollback()
		}
	}()
	call(dbc)
	dbc.Commit()
}
