package gutils

import (
	"reflect"
	"strconv"
)

func SaveValByInterface(v reflect.Value, obj interface{}) {
	t := reflect.TypeOf(obj)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64:
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int64:
			v.Set(reflect.ValueOf(obj))
		case reflect.String:
			v2, _ := strconv.ParseInt(obj.(string), 10, 64)
			v.SetInt(v2)
		default:
			Error(500, "数据类型错误")
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch t.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v.Set(reflect.ValueOf(obj))
		case reflect.String:
			v2, _ := strconv.ParseUint(obj.(string), 10, 64)
			v.SetUint(v2)
		default:
			Error(500, "数据类型错误")
		}
	case reflect.Float32, reflect.Float64:
		switch t.Kind() {
		case reflect.Float32, reflect.Float64:
			v.Set(reflect.ValueOf(obj))
		case reflect.String:
			v2, _ := strconv.ParseFloat(obj.(string), 64)
			v.SetFloat(v2)
		default:
			Error(500, "数据类型错误")
		}
	case reflect.String:
		switch t.Kind() {
		case reflect.String:
			v.SetString(obj.(string))
		default:
			Error(500, "数据类型错误")
		}
	case reflect.Slice:
		switch t.Kind() {
		case reflect.Slice:
			v.Set(reflect.ValueOf(obj))
		default:
			Error(500, "数据类型错误")
		}
	case reflect.Map:
		switch t.Kind() {
		case reflect.Map:
			v.Set(reflect.ValueOf(obj))
		default:
			Error(500, "数据类型错误")
		}
	case reflect.Struct:
		switch t.Kind() {
		case reflect.Struct:
			if v.Kind().String() != t.Kind().String() {
				Error(500, "数据类型错误")
			}
			v.Set(reflect.ValueOf(obj))
		default:
			Error(500, "数据类型错误")
		}
	}
}
