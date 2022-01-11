package es_driver

import (
	"errors"
	"github.com/ghf-go/nannan/gutils"
	"reflect"
)

type esMgetResonse struct {
	Docs []struct {
		Index       string                 `json:"_index"`
		Type        string                 `json:"_type"`
		Id          string                 `json:"_id"`
		Version     int                    `json:"_version"`
		SeqNo       int                    `json:"_seq_no"`
		PrimaryTerm int                    `json:"_primary_term"`
		Found       bool                   `json:"found"`
		Source      map[string]interface{} `json:"_source"`
	} `json:"docs"`
}

func (r *esMgetResonse) saveObj(obj interface{}) error {
	if len(r.Docs) == 0 {
		return nil
	}
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Map {
		return errors.New("mGet 返回必须是map")
	}
	itemType := t.Elem()
	m := reflect.ValueOf(obj)
	for _, item := range r.Docs {
		if !item.Found {
			continue
		}
		key := reflect.New(t.Key()).Elem()
		gutils.SaveValByInterface(key, item.Id)
		//fmt.Printf("类型 %s -> %s \n", t.String(), itemType.String())
		switch itemType.Kind() {
		case reflect.Interface:
			v := map[string]interface{}{}
			vv := reflect.ValueOf(v)
			vv.SetMapIndex(reflect.ValueOf("_id"), reflect.ValueOf(item.Id))
			for k, v3 := range item.Source {
				vv.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v3))
			}
			m.SetMapIndex(key, vv)

		case reflect.Struct:
			v := reflect.New(itemType).Elem()
			t := v.Type()
			mapId := map[string]int{}
			for i := 0; i < t.NumField(); i++ {
				tn := t.Field(i).Tag.Get("es_field")
				if tn == "" {
					continue
				}
				mapId[tn] = i
			}

			if i, ok := mapId["_id"]; ok {
				gutils.SaveValByInterface(v.Field(i), item.Id)
			}

			for k, vv := range item.Source {
				if i, ok := mapId[k]; ok {
					gutils.SaveValByInterface(v.Field(i), vv)
				}
			}
			m.SetMapIndex(key, v)
		default:

			return errors.New("mGet 返回必须是map")

		}
	}
	return nil
}
