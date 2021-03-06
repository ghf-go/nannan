package web

import (
	"github.com/ghf-go/nannan/gutils"
	"github.com/ghf-go/nannan/mod"
	"net/http"
	"strings"
	"time"
)

type Router struct {
	data map[string]*RouterGroup
}

var (
	_router = &Router{data: map[string]*RouterGroup{}}
)

//注册静态文件路径
func RegisterStaticDir(path, dir string) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	_newHandle.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

//注册分组
func RegisterRouterGroup(prefix string, funcName func(group *RouterGroup)) {
	_router.RegisterRouterGroup(prefix, funcName)
}

//注册群组
func (r *Router) RegisterRouterGroup(prefix string, funcName func(group *RouterGroup)) {
	group := &RouterGroup{
		data: map[string]map[string]func(ctx *EngineCtx) error{},
	}
	funcName(group)
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	r.data[prefix] = group
}
func error404(engineCtx *EngineCtx) {
	engineCtx.ReturnStatusCode(404)
	engineCtx.Write([]byte("页面不存在"))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	st := time.Now()
	engineCtx := &EngineCtx{
		gresponse: &gresponse{},
		ReqID:     time.Now().UnixNano(),
		Req:       req,
		rep:       w,
	}
	runMiddleWare(engineCtx, func(ctx *EngineCtx) {
		path := ctx.Req.URL.Path[1:]
		arr := strings.Split(path, "/")
		gdir := "/"
		for i, v := range arr {
			gdir += v + "/"
			if group, ok := r.data[gdir]; ok {
				engineCtx.GroupPath = gdir
				engineCtx.NodePath = strings.Join(arr[i+1:], "/")
				//mod.Debug("程序分组 %s  -> %s",gdir ,engineCtx.NodePath)
				group.run(engineCtx)
				return
			}
		}
		error404(engineCtx)
	})
	e := time.Now()
	mod.Debug("[%s] %dms %s", gutils.FormatTime(e), e.Sub(st).Microseconds(), req.URL.Path)
}
