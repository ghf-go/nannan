package web

type Router struct {
	data map[string]*RouterGroup
}
//注册群组
func (r *Router) RegisterRouterGroup(prefix string,funcName func(group *RouterGroup))  {
	group := &RouterGroup{}
	funcName(group)
	r.data[prefix] = group

}
