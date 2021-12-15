package web

import (
	"bytes"
	"net/http"
)

type gresponse struct {
	header http.Header
	cookies []*http.Cookie
	httpCode int
	buf *bytes.Buffer
}
func (r *gresponse) Header() http.Header{
	if r.header == nil{
		r.header = http.Header{}
	}
	return r.header
}
func (r *gresponse) SetCookie(cookie *http.Cookie)  {
	if r.cookies == nil{
		r.cookies = []*http.Cookie{cookie}
	}else{
		r.cookies = append(r.cookies,cookie)
	}
}
func (r *gresponse) ReturnStatusCode(code int)  {
	r.httpCode = code
}
func (r *gresponse) Write(b []byte)  {
	if r.buf == nil{
		r.buf = &bytes.Buffer{}
	}
	r.buf.Write(b)
}
