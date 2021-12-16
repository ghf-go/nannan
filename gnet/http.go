package gnet

import "net/http"

type Ghttp string

func (h Ghttp) Get(query string) (resp *http.Response, err error) {
	 return  http.Get(string(h) + query)
}

func (h Ghttp) Put()  {

}
func (h Ghttp) Post()  {

}