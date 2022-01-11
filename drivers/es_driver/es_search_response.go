package es_driver

import (
	"reflect"
)

type esSearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string                 `json:"_index"`
			Type   string                 `json:"_type"`
			Id     string                 `json:"_id"`
			Score  float64                `json:"_score"`
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (r *esSearchResponse) saveObj(obj interface{}) error {
	t := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch t.Kind() {
	case reflect.Map:
		objV.SetMapIndex(reflect.ValueOf("total"), reflect.ValueOf(r.Hits.Total.Value))
		list := []map[string]interface{}{}
		for _, item := range r.Hits.Hits {
			val := map[string]interface{}{
				"_id": item.Id,
			}
			for k, v := range item.Source {
				val[k] = v
			}
			list = append(list, val)
		}
		objV.SetMapIndex(reflect.ValueOf("list"), reflect.ValueOf(list))
	}

	return nil
}
