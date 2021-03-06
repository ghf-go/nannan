package web

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/mod"
	"reflect"
	"strings"
)

type RouterGroup struct {
	data map[string]map[string]func(ctx *EngineCtx) error
}

func (r *RouterGroup) resetPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	return path
}
func (r *RouterGroup) POST(path string, funcName func(ctx *EngineCtx) error) {
	if _, ok := r.data["POST"]; !ok {
		r.data["POST"] = map[string]func(ctx *EngineCtx) error{
			r.resetPath(path): funcName,
		}
	} else {
		r.data["POST"][r.resetPath(path)] = funcName
	}
}
func (r *RouterGroup) GET(path string, funcName func(ctx *EngineCtx) error) {
	if _, ok := r.data["GET"]; !ok {
		r.data["GET"] = map[string]func(ctx *EngineCtx) error{
			r.resetPath(path): funcName,
		}
	} else {
		r.data["GET"][r.resetPath(path)] = funcName
	}
}
func (r *RouterGroup) PUT(path string, funcName func(ctx *EngineCtx) error) {
	if _, ok := r.data["PUT"]; !ok {
		r.data["PUT"] = map[string]func(ctx *EngineCtx) error{
			r.resetPath(path): funcName,
		}
	} else {
		r.data["PUT"][r.resetPath(path)] = funcName
	}
}
func (r *RouterGroup) DELETE(path string, funcName func(ctx *EngineCtx) error) {
	if _, ok := r.data["DELETE"]; !ok {
		r.data["DELETE"] = map[string]func(ctx *EngineCtx) error{
			r.resetPath(path): funcName,
		}
	} else {
		r.data["DELETE"][r.resetPath(path)] = funcName
	}
}
func (r *RouterGroup) ANY(path string, funcName func(ctx *EngineCtx) error) {
	if _, ok := r.data["ANY"]; !ok {
		r.data["ANY"] = map[string]func(ctx *EngineCtx) error{
			r.resetPath(path): funcName,
		}
	} else {
		r.data["ANY"][r.resetPath(path)] = funcName
	}
}
func (r *RouterGroup) run(engineCtx *EngineCtx) {
	defer func() {
		if e := recover(); e != nil {
			if gutils.IsError(e) {
				e2 := e.(gutils.BaseErr)
				mod.Error("系统错误 %d %s", e2.Code, e2.Msg)
				engineCtx.JsonFail(e2.Code, e2.Msg)
			} else if reflect.TypeOf(e).Kind() == reflect.String {
				mod.Error("系统错误 string %s", e)
				engineCtx.JsonFail(500, e.(string))
			} else {
				mod.Error("系统错误 error %s", e.(error).Error())
				engineCtx.JsonFail(500, e.(error).Error())
			}
		}
	}()
	if urls, ok := r.data[engineCtx.Req.Method]; ok {
		if f, o := urls[engineCtx.NodePath]; o {
			f(engineCtx)
			return
		}
	}
	if urls, ok := r.data["ANY"]; ok {
		if f, o := urls[engineCtx.NodePath]; o {
			f(engineCtx)
			return
		}
	}
	error404(engineCtx)
}
