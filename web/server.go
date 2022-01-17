package web

import (
	"fmt"
	"github.com/ghf-go/nannan/mod"
	"log"
	"net/http"
)

var (
	_newHandle = &http.ServeMux{}
)

func WebStart() {
	_newHandle.Handle("/", _router)
	conf := mod.GetConf("app.webport")
	server := &http.Server{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler:  _newHandle,
		ErrorLog: log.Default(),
	}
	server.ListenAndServe()
}
