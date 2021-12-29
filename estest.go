package main

import (
	"github.com/ghf-go/nannan/db"
	"os"
)

type RUser struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}
type User struct {
	*db.EsResponseBase
	*db.EsResponseBaseDoc
	Data *RUser `json:"_source"`
}
type UserMget struct {
	*db.EsResponseBase
	Docs []struct {
		*db.EsResponseDocsBaseDoc
		Data RUser `json:"_source"`
	} `json:"docs"`
}

func main() {
	os.Setenv("es.test", "mem://dev_gay/us.ggvjj.ml:9200")
	es := db.GetEsClient("test")
	//r, e := es.Update("user", "65431", map[string]interface{}{
	//	"user_id": 65431,
	//	"name":    "测试2",
	//})
	//if e != nil {
	//	glog.Error("插入失败1 %s", e.Error())
	//} else if !r.IsSuccess() {
	//	glog.Error("插入失败2 %s %v", r, r)
	//} else {
	//	glog.Debug("创建成功 %s, %v", r, r)
	//}
	//u := &User{}
	//e := es.Find("user", "65431", u)
	//if e != nil {
	//	glog.Error("查找失败 %s", e.Error())
	//} else {
	//	glog.Debug("user id:%s userid:%d name %s", u.Id, u.Data.UserId, u.Data.Name)
	//}
	//r, e := es.Delete("user", "65431")
	//if e != nil {
	//	glog.Error("删除失败 %s", e.Error())
	//} else if !r.IsSuccess() {
	//	glog.Error("删除失败 %s %s", r, r)
	//} else {
	//	glog.Debug("删除成功")
	//}
	//o := &UserMget{}
	//e := es.MGet("user", o, "65431", "FnWf6n0Be21VpBLDHy3L")
	//if e != nil {
	//	glog.Error("查找失败 %s", e.Error())
	//} else {
	//	glog.Debug("查找完成 %s %v", o.Docs, o.Docs)
	//}

}
