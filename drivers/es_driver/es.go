package es_driver

import (
	"bytes"
	"encoding/json"
	"github.com/ghf-go/nannan/drivers"
	"io"
	"io/ioutil"
	"net/http"
)

type EsClient struct {
	Hosts     []string
	DbName    string
	ReqIndex  int
	HostCount int
}

func (es *EsClient) getHost() string {
	if es.HostCount == 0 {
		es.HostCount = len(es.Hosts)
	}
	es.ReqIndex += 1
	return "http://" + es.Hosts[es.ReqIndex%es.HostCount] + "/" + es.DbName + "/"
}

//添加记录
func (es *EsClient) Create(tbName string, obj interface{}) (*EsResponseUpdate, error) {
	ret := &EsResponseUpdate{}
	return ret, es.do(http.MethodPost, es.getHost()+tbName, obj, ret)
}

//更新记录
func (es *EsClient) Update(tbName, id string, obj interface{}) (*EsResponseUpdate, error) {
	ret := &EsResponseUpdate{}
	return ret, es.do(http.MethodPost, es.getHost()+tbName+"/"+id, obj, ret)
}

// 删除记录
func (es *EsClient) Delete(tbName, id string) (*EsResponseUpdate, error) {
	ret := &EsResponseUpdate{}
	return ret, es.do(http.MethodDelete, es.getHost()+tbName+"/"+id, nil, ret)
}
func (es *EsClient) do(method, url string, body interface{}, obj interface{}) error {
	var bb io.Reader
	if body != nil {
		bb1, e := json.Marshal(body)
		if e != nil {
			drivers.Error("ES 提交内容格式化 %s-> %s error:%s", method, url, e.Error())
			return e
		}
		bb = bytes.NewReader(bb1)
	}

	req, e := http.NewRequest(method, url, bb)
	if e != nil {
		drivers.Error("ES 创建请求结构体 %s-> %s error:%s", method, url, e.Error())
		return e
	}
	req.Header.Add("Content-Type", "application/json")
	r, e := http.DefaultClient.Do(req)
	if e != nil {
		drivers.Error("ES 发送数据 %s-> %s error:%s", method, url, e.Error())
		return e
	}

	defer r.Body.Close()
	buf, e := ioutil.ReadAll(r.Body)
	if e != nil {
		drivers.Error("ES 读取返回 %s-> %s error:%s", method, url, e.Error())
		return e
	}

	e = json.Unmarshal(buf, obj)
	if e != nil {
		drivers.Error("ES 结果转换 %s-> %s error:", method, url, e.Error())
		return e
	}

	return nil
}

//查询一条记录
func (es *EsClient) Find(tbName, id string, obj interface{}) error {
	return es.do(http.MethodGet, es.getHost()+tbName+"/"+id, nil, obj)
}

//批量查询
func (es *EsClient) MGet(tbName string, obj interface{}, ids ...string) error {
	return es.do(http.MethodPost, es.getHost()+tbName+"/_mget", map[string]interface{}{"ids": ids}, obj)
}
func (es *EsClient) NewQuery(tbName string) *esQeury {
	return newEsQuery(es, tbName)
}
