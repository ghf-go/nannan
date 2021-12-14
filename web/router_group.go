package web

import (
	"net/http"
	"strings"
	"time"
)

type RouterGroup struct {
	data map[string]map[string]func(ctx *EngineCtx) error
}

func (r *RouterGroup) resetPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	//if strings.HasSuffix(path, "/") {
	//	path = path[:-1]
	//}
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
func (r *RouterGroup) run(w http.ResponseWriter, req *http.Request, i int, args []string) {
	if urls, ok := r.data[req.Method]; ok {
		al := len(args)
		if i+1 < al {
			path := strings.Join(args[i+1:], "/")
			if f, o := urls[path]; o {
				defer func() {
					if e := recover();e != nil{
						error404(w, req)
					}
				}()
				f(&EngineCtx{
					ReqID: time.Now().UnixNano(),
					Req:   req,
					Rep:   w,
				})
			}
		}
	}
	error404(w, req)
}
