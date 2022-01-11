package es_driver

import (
	"errors"
	"github.com/ghf-go/nannan/gutils"
	"reflect"
)

type esFindResonse struct {
	Index       string                 `json:"_index"`
	Type        string                 `json:"_type"`
	Id          string                 `json:"_id"`
	Version     int                    `json:"_version"`
	SeqNo       int                    `json:"_seq_no"`
	PrimaryTerm int                    `json:"_primary_term"`
	Found       bool                   `json:"found"`
	Source      map[string]interface{} `json:"_source"`
}

func (r *esFindResonse) saveObj(obj interface{}) error {
	if !r.Found {
		return errors.New("没有查到记录")
	}
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Map {
		vv := reflect.ValueOf(obj)
		vv.SetMapIndex(reflect.ValueOf("_id"), reflect.ValueOf(r.Id))
		for k, v3 := range r.Source {
			vv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v3))
		}
		return nil
	}
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return errors.New("数据必须是引用结构体")
	}
	mapId := map[string]int{}
	for i := 0; i < t.Elem().NumField(); i++ {
		tn := t.Elem().Field(i).Tag.Get("es_field")
		if tn == "" {
			continue
		}
		mapId[tn] = i
	}
	v := reflect.ValueOf(obj).Elem()
	if i, ok := mapId["_id"]; ok {
		gutils.SaveValByInterface(v.Field(i), r.Id)
	}

	for k, vv := range r.Source {
		if i, ok := mapId[k]; ok {
			gutils.SaveValByInterface(v.Field(i), vv)
		}
	}
	return nil
}
