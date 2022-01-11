package test

import (
	"github.com/ghf-go/nannan/def"
	"github.com/ghf-go/nannan/mod"
)

func esSave(args []string) {
	mod.Debug("开始执行")
	for i := 0; i < 100; i++ {
		es := mod.GetEsClient("es")
		es.Update("users", "1234", def.Data{"name": "张三"})
		//mod.Debug("se-save %v", es.Update("users", "1234", def.Data{"name": "张三"}))
	}
	mod.Debug("结束执行")
}
func esDelete(args []string) {
	mod.Debug("开始执行")
	for i := 0; i < 100; i++ {
		es := mod.GetEsClient("es")
		es.Delete("users", "1234")
		//mod.Debug("se-delete %v ", es.Delete("users", "1234"))
	}
	mod.Debug("结束执行")
}

type User struct {
	Id   string `es_field:"_id"`
	Name string `es_field:"name"`
}

func esFind(args []string) {
	mod.Debug("开始执行")
	for i := 0; i < 10; i++ {
		es := mod.GetEsClient("es")
		obj := &User{}
		//es.Find("users", "1234", obj)
		mod.Debug("se-find %v %v", es.Find("users", "1234", obj), obj)
	}
	mod.Debug("结束执行")
}
func esMget(args []string) {
	mod.Debug("开始执行")
	for i := 0; i < 10; i++ {
		es := mod.GetEsClient("es")
		data := map[string]User{}
		e := es.MGet("users", data, "1234")
		mod.Debug("se-mget %v %v", e, data)
	}
	mod.Debug("结束执行")
}
func esSearch(args []string) {
	mod.Debug("开始执行")
	for i := 0; i < 10; i++ {
		data := map[string]interface{}{}
		e := mod.GetEsClient("es").NewQuery("users").Query(data)
		mod.Debug("se-search %v %v", e, data)
	}
	mod.Debug("结束执行")
}
