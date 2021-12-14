package web

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

type EngineCtx struct {
	ReqID int64
	Req   *http.Request
	Rep   http.ResponseWriter
	ip string
}
// 输出JSON
func (engine *EngineCtx) json(code int, msg string, data interface{}) error {
	ret := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	b, e := json.Marshal(ret)
	if e == nil {
		engine.Rep.Write(b)
	}
	return e
}
//输出错误的json
func (engine *EngineCtx) JsonFail(code int, msg string) error {
	return engine.json(code, msg, nil)
}
//输出正确的json
func (engine *EngineCtx) JsonSuccess(data interface{}) error {
	return engine.json(0, "", data)
}
//显示网页
func (engine *EngineCtx) Display(tpl string,data interface{}) error{
	return nil
}
// 获取IP
func (engine *EngineCtx) GetIP()string{
	if engine.ip != ""{
		return engine.ip
	}
	xForwardedFor := engine.Req.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(engine.Req.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(engine.Req.RemoteAddr)); err == nil {
		return ip
	}

	return "0.0.0.0"
}
