package web

import "net/http"

var (
	_newHandle = &http.ServeMux{}
)

func WebStart(addr string)  {
	_newHandle.Handle("/",_router)
	server := &http.Server{
		Addr: addr,
		Handler: _newHandle,
	}
	server.ListenAndServe()
}
