package web

import (
	"net/http"
	"os"
)

var (
	_newHandle = &http.ServeMux{}
)

func WebStart() {
	_newHandle.Handle("/", _router)
	server := &http.Server{
		Addr:    os.Getenv("app.web"),
		Handler: _newHandle,
	}
	server.ListenAndServe()
}
