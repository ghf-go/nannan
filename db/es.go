package db

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type EsClient struct {
	hosts  []string
	dbName string
}

func (es *EsClient) getHost() string {
	return ""
}

//添加记录
func (es *EsClient) Create(tbName string, obj interface{}) (*EsResponse, error) {
	return es.do(http.MethodPost, es.getHost()+tbName, obj, nil)
}

//更新记录
func (es *EsClient) Update(tbName string, obj interface{}) (*EsResponse, error) {
	return es.do(http.MethodPut, es.getHost()+tbName, obj, nil)
}

// 删除记录
func (es *EsClient) Delete(tbName, id string) (*EsResponse, error) {
	return es.do(http.MethodDelete, es.getHost()+tbName+"/"+id, nil, nil)
}
func (es *EsClient) do(method, url string, body interface{}, obj interface{}) (*EsResponse, error) {
	var bb io.Reader
	if obj != nil {
		bb1, e := json.Marshal(body)
		if e != nil {
			return nil, e
		}
		bb = bytes.NewReader(bb1)
	}

	req, e := http.NewRequest(method, url, bb)
	if e != nil {
		return nil, e
	}
	res := &EsResponse{}
	if obj != nil {
		res.Source = obj
	}
	r, e := http.DefaultClient.Do(req)
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()
	buf, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return nil, e
	}

	e = json.Unmarshal(buf, res)
	if e != nil {
		return nil, e
	}
	res.HttpCode = r.StatusCode
	return res, e
}

//查询一条记录
func (es *EsClient) Find(tbName, id string) (*EsResponse, error) {
	return es.do(http.MethodGet, es.getHost()+tbName+"/"+id, nil, nil)
}

//查询一条记录
func (es *EsClient) FindObj(tbName, id string, obj interface{}) error {
	_, e := es.do(http.MethodGet, es.getHost()+tbName+"/"+id, nil, obj)
	if e != nil {
		return e
	}
	return nil
}

//批量查询
func (es *EsClient) MGet(tbName string, ids ...string) (*EsResponse, error) {
	return es.do(http.MethodPost, es.getHost()+tbName+"/_mget", map[string]interface{}{"ids": ids}, nil)
}
func (es *EsClient) NewQuery(tbName string) *esQeury {
	return newEsQuery(es, tbName)
}
