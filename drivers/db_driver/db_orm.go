package db_driver

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ghf-go/nannan/mod"
	"reflect"
	"strings"
	"time"
)

var (
	NOT_ROW = errors.New("没有记录")
)

const (
	_cokumn_name = "column"
)

func saveObj(rows *sql.Rows, obj interface{}) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		mod.Error("错误 保存对象必须是指针对象 %s", t.Kind())
		return errors.New("保存对象必须是指针对象")
	}
	t = t.Elem()
	val := reflect.ValueOf(obj).Elem()
	columns, e := rows.Columns()
	if e != nil {
		mod.Error("错误 %s", e.Error())
		return e
	}
	fm := _getColumMapByType(columns, t)
	ln := len(columns)
	defer rows.Close()
	if rows.Next() {
		e := _saveRow(rows, ln, fm, t, val)
		if e != nil {
			return e
		}
		return nil
	} else {
		return errors.New("保存对象必须是指针对象")
	}
}
func saveObjList(rows *sql.Rows, obj interface{}) error {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Slice {
		return errors.New("保存对象必须是指针数组")
	}
	isRef := false
	var tarObj reflect.Type
	if t.Elem().Elem().Kind() == reflect.Ptr {
		isRef = true
		tarObj = t.Elem().Elem().Elem()
	} else {
		tarObj = t.Elem().Elem()
	}
	columns, e := rows.Columns()
	if e != nil {
		return e
	}

	fm := _getColumMapByType(columns, tarObj)
	ln := len(columns)
	defer rows.Close()
	ret := []reflect.Value{}
	for rows.Next() {
		val := reflect.New(tarObj)
		e := _saveRow(rows, ln, fm, tarObj, val)
		if e != nil {

			return e
		}
		if isRef {
			ret = append(ret, val)
		} else {
			ret = append(ret, val.Elem())
		}

	}
	a2 := reflect.Append(reflect.ValueOf(obj).Elem(), ret...)
	reflect.ValueOf(obj).Elem().Set(a2)
	return nil
}

//获取对象的字段名字列表
func getObjFields(obj interface{}) []string {
	ret := []string{}
	t := reflect.TypeOf(obj)
	for {
		switch t.Kind() {
		case reflect.Ptr:
			t = t.Elem()
		case reflect.Map:
			t = t.Elem()
		case reflect.Slice:
			t = t.Elem()
		case reflect.Struct:
			break

		}
	}
	fnum := t.NumField()
	for i := 0; i < fnum; i++ {
		tag := t.Field(i).Tag.Get(_cokumn_name)
		if tag != "" {
			ret = append(ret, tag)
		}
	}
	return ret
}
func getObjFieldStr(obj interface{}) string {
	ret := []string{}
	t := reflect.TypeOf(obj)
	isBreak := true
	for isBreak {
		switch t.Kind() {
		case reflect.Ptr:
			t = t.Elem()
		case reflect.Map:
			t = t.Elem()
		case reflect.Slice:
			t = t.Elem()
		case reflect.Struct:
			isBreak = false
			break
		}
	}

	fnum := t.NumField()
	for i := 0; i < fnum; i++ {
		tag := t.Field(i).Tag.Get(_cokumn_name)
		if tag != "" {
			ret = append(ret, tag)
		}
	}
	return fmt.Sprintf("`%s`", strings.Join(ret, "`,`"))
}

//保存数据
func _saveRow(rows *sql.Rows, cl int, fm map[int]int, t reflect.Type, obj reflect.Value) error {
	args := []interface{}{}
	for i := 0; i < cl; i++ {
		if indx, ok := fm[i]; ok {
			if strings.HasSuffix(t.Field(indx).Type.String(), "time.Time") {
				var iv time.Time
				args = append(args, &iv)
			} else {
				switch t.Field(indx).Type.Kind() {
				case reflect.Int:
					var vv int
					args = append(args, &vv)
				case reflect.Int8:
					var vv int8
					args = append(args, &vv)
				case reflect.Int16:
					var vv int16
					args = append(args, &vv)
				case reflect.Int32:
					var vv int32
					args = append(args, &vv)
				case reflect.Int64:
					var vv int64
					args = append(args, &vv)
				case reflect.Uint:
					var vv uint
					args = append(args, &vv)
				case reflect.Uint8:
					var vv uint
					args = append(args, &vv)
				case reflect.Uint16:
					var vv uint16
					args = append(args, &vv)
				case reflect.Uint32:
					var vv uint32
					args = append(args, &vv)
				case reflect.Uint64:
					var vv uint64
					args = append(args, &vv)
				case reflect.Float32:
					var vv float32
					args = append(args, &vv)
				case reflect.Float64:
					var vv float64
					args = append(args, &vv)
				case reflect.String:
					var vv string
					args = append(args, &vv)
				}
			}
		} else {
			var vv interface{}
			args = append(args, &vv)
		}
	}
	e := rows.Scan(args...)
	if e != nil {
		mod.Error("cuowu %s", e.Error())
		return e
	}
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}
	for ci, i := range fm {
		obj.Field(i).Set(reflect.ValueOf(args[ci]).Elem())
	}
	return nil
}
func _getColumMapByType(columns []string, t reflect.Type) map[int]int {
	ret := map[int]int{}
	fn := t.NumField()
	for i := 0; i < fn; i++ {
		tag := t.Field(i).Tag.Get(_cokumn_name)
		for ind, nm := range columns {
			if nm == tag {
				ret[ind] = i
			}
		}
	}
	return ret
}
