package web

type RouterGroup struct {
	data map[string]map[string]func(ctx EngineCtx)error
}

func (r *RouterGroup) POST(path string ,funcName func(ctx EngineCtx)error) {
	if _,ok:=r.data["POST"];!ok {
		r.data["POST"] = map[string]func(ctx EngineCtx) error{
			path: funcName,
		}
	}else{
		r.data["POST"][path] = funcName
	}
}
func (r *RouterGroup) GET(path string ,funcName func(ctx EngineCtx)error)  {
	if _,ok:=r.data["GET"];!ok {
		r.data["GET"] = map[string]func(ctx EngineCtx) error{
			path: funcName,
		}
	}else{
		r.data["GET"][path] = funcName
	}
}
func (r *RouterGroup) PUT(path string ,funcName func(ctx EngineCtx)error)  {
	if _,ok:=r.data["PUT"];!ok {
		r.data["PUT"] = map[string]func(ctx EngineCtx) error{
			path: funcName,
		}
	}else{
		r.data["PUT"][path] = funcName
	}
}
func (r *RouterGroup) DELETE(path string ,funcName func(ctx EngineCtx)error)  {
	if _,ok:=r.data["DELETE"];!ok {
		r.data["DELETE"] = map[string]func(ctx EngineCtx) error{
			path: funcName,
		}
	}else{
		r.data["DELETE"][path] = funcName
	}
}
func (r *RouterGroup) ANY(path string ,funcName func(ctx EngineCtx)error)  {
	if _,ok:=r.data["ANY"];!ok {
		r.data["ANY"] = map[string]func(ctx EngineCtx) error{
			path: funcName,
		}
	}else{
		r.data["ANY"][path] = funcName
	}
}