package test

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/mod"
)

func esSave(args []string) {
	es := mod.GetEsClient("es")

	mod.Debug("se-save %v", es.Update("users", "1234", def.Data{"name": "张三"}))
}
func esDelete(args []string) {
	es := mod.GetEsClient("es")
	mod.Debug("se-delete %v ", es.Delete("users", "1234"))
}

type User struct {
	Id   string `es_field:"_id"`
	Name string `es_field:"name"`
}

func esFind(args []string) {
	es := mod.GetEsClient("es")
	obj := &User{}
	mod.Debug("se-find %v %v", es.Find("users", "1234", obj), obj)
}
func esMget(args []string) {
	es := mod.GetEsClient("es")
	data := map[string]User{}
	e := es.MGet("users", data, "1234")
	mod.Debug("se-mget %v %v", e, data)
}
func esSearch(args []string) {
	data := map[string]interface{}{}
	e := mod.GetEsClient("es").NewQuery("users").Query(data)
	mod.Debug("se-search %v %v", e, data)
}
