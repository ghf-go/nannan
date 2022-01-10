package db_driver

import (
	"database/sql"
	"fmt"
	"github.com/ghf-go/nannan/mod"
	"strings"
)

type Query struct {
	fields    string
	table     *Table
	where     string
	onsql     string
	start     int
	pagesize  int
	orderList []string
	joinTable []string
	whereArgs []interface{}
	onArgs    []interface{}
}

func (q *Query) Join(table string) *Query {
	if q.joinTable == nil {
		q.joinTable = []string{}
	}
	q.joinTable = append(q.joinTable, fmt.Sprintf("JOIN `%s`", table))
	return q
}
func (q *Query) JoinLeft(table string) *Query {
	if q.joinTable == nil {
		q.joinTable = []string{}
	}
	q.joinTable = append(q.joinTable, fmt.Sprintf("JOIN LEFT `%s`", table))
	return q
}
func (q *Query) JoinRight(table string) *Query {
	if q.joinTable == nil {
		q.joinTable = []string{}
	}
	q.joinTable = append(q.joinTable, fmt.Sprintf("JOIN RIGHT `%s`", table))
	return q
}
func (q *Query) On(sql string, args ...interface{}) *Query {
	q.onArgs = append(q.onArgs, args...)
	if q.onsql == "" {
		q.onsql = sql
	} else {
		q.onsql += " AND " + sql
	}
	return q
}
func (q *Query) Select(fileds []string) *Query {
	q.SelectRaw(fmt.Sprintf("`%s`", strings.Join(fileds, "`,`")))
	return q
}
func (q *Query) SelectRaw(fileds string) *Query {
	if q.fields == "" {
		q.fields = fileds
	} else {
		q.fields += "," + fileds
	}
	return q
}
func (q *Query) Where(sql string, args ...interface{}) *Query {
	if q.whereArgs == nil {
		q.whereArgs = []interface{}{}
	}
	q.whereArgs = append(q.whereArgs, args...)
	if q.where == "" {
		q.where = sql
	} else {
		q.where += " AND " + sql
	}
	return q
}
func (q *Query) Count(field string) int64 {
	return q._queryNum(field, "count")
}
func (q *Query) Sum(field string) int64 {
	return q._queryNum(field, "sum")
}
func (q *Query) Avg(field string) int64 {
	return q._queryNum(field, "avg")
}
func (q *Query) Order(name, order string) *Query {
	if q.orderList == nil {
		q.orderList = []string{}
	}
	q.orderList = append(q.orderList, fmt.Sprintf("`%s` %s", name, order))
	return q
}
func (q *Query) Frist(obj interface{}) error {
	r, e := q._fetch(getObjFieldStr(obj))
	if e != nil {
		mod.Error("sql 错误 %s", e.Error())
		return e
	}
	return saveObj(r, obj)
}
func (q *Query) FetchAll(obj interface{}) error {
	r, e := q._fetch(getObjFieldStr(obj))
	if e != nil {
		return e
	}
	return saveObjList(r, obj)
}

func (q *Query) Skip(start int) *Query {
	q.start = start
	return q
}
func (q *Query) Limit(pagesize int) *Query {
	q.pagesize = pagesize
	return q
}

func (q *Query) UpdateMap(data Data) int64 {
	sets := []string{}
	args := []interface{}{}
	for k, v := range data {
		if strings.HasSuffix(k, "+") {
			sets = append(sets, fmt.Sprintf("`%s` = `%s` + ?", k, k))
			args = append(args, v)
		} else if strings.HasSuffix(k, "-") {
			sets = append(sets, fmt.Sprintf("`%s` = `%s` - ?", k, k))
			args = append(args, v)
		} else if strings.HasSuffix(k, "*") {
			sets = append(sets, fmt.Sprintf("`%s` = `%s` * ?", k, k))
			args = append(args, v)
		} else if strings.HasSuffix(k, "/") {
			sets = append(sets, fmt.Sprintf("`%s` = `%s` / ?", k, k))
			args = append(args, v)
		} else {
			sets = append(sets, fmt.Sprintf("`%s` =  ?", k))
			args = append(args, v)
		}
	}
	sql := fmt.Sprintf("UPDATE `%s` SET %s", q.table.table, strings.Join(sets, ","))
	if q.where != "" {
		sql += " WHERE " + q.where
		args = append(args, q.whereArgs...)
	}
	if len(q.orderList) > 0 {
		sql += " ORDER BY " + strings.Join(q.orderList, ",")
	}
	if q.pagesize > 0 {
		sql += fmt.Sprintf(" LIMIT %d", q.pagesize)
		if q.start > 0 {
			sql += fmt.Sprintf(" OFFSET %s", q.start)
		}
	}
	return q.table.UpdateSql(sql, args...)
}

func (q *Query) Reset() *Query {
	return q
}

func (q *Query) _fetch(fields string) (*sql.Rows, error) {
	sql := ""
	args := []interface{}{}
	sql = fmt.Sprintf("SELECT %s FROM `%s` ", fields, q.table.table)
	if len(q.joinTable) > 0 {
		sql += " " + strings.Join(q.joinTable, " ")
	}
	if q.onsql != "" {
		sql += " ON " + q.onsql
		args = append(args, q.onArgs...)
	}
	if q.where != "" {
		sql += " WHERE " + q.where
		args = append(args, q.whereArgs...)
	}
	if len(q.orderList) > 0 {
		sql += " ORDER BY " + strings.Join(q.orderList, ",")
	}
	if q.pagesize > 0 {
		sql += fmt.Sprintf(" LIMIT %d", q.pagesize)
		if q.start > 0 {
			sql += fmt.Sprintf(" OFFSET %d", q.start)
		}
	}

	return q.table.Query(sql, args...)
}

//删除数据
func (q *Query) Delete() int64 {
	sql := fmt.Sprintf("DELETE FROM `%s`", q.table.TableName())
	args := []interface{}{}
	if q.where != "" {
		sql += " WHERE " + q.where
		args = append(args, q.whereArgs...)
	}
	if len(q.orderList) > 0 {
		sql += " ORDER BY " + strings.Join(q.orderList, ",")
	}
	if q.pagesize > 0 {
		sql += fmt.Sprintf(" LIMIT %d", q.pagesize)
		if q.start > 0 {
			sql += fmt.Sprintf(" OFFSET %d", q.start)
		}
	}
	return q.table.DeleteSql(sql, args...)

}
func (q *Query) _queryNum(f, funcname string) int64 {
	sql := ""
	args := []interface{}{}
	sql = fmt.Sprintf("SELECT %s(`%s`) FROM `%s` ", strings.ToUpper(funcname), f, q.table.table)
	if len(q.joinTable) > 0 {
		sql += " " + strings.Join(q.joinTable, " ")
	}
	if q.onsql != "" {
		sql += " ON " + q.onsql
		args = append(args, q.onArgs...)
	}
	if q.where != "" {
		sql += " WHERE " + q.where
		args = append(args, q.whereArgs...)
	}
	return q.table.QueryNum(sql, args...)
}
