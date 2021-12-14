# nannan

# 配置
## 全局配置
### web服务
> os.Setenv("app.web",":80") //监听80端口
## 缓存设置
### memcache 配置
> os.Setenv("key","memcache://")
### redis 配置
> os.Setenv("key","redis://user:passwd@ip:port/?retries=3&db=1")
### redis 群集配置
> os.Setenv("key","redis_cluster://"))
### redis 哨兵配置
> os.Setenv("key","redis_sentinel://user:passwd@ip:port/?retries=3&db=1&servers=ip:port,ip:port"))

# 数据库使用

# 路由
```go
web.RegisterRouterGroup("/admin",func (group *RouterGroup){
	group.POST("abc",function(x EngineCtx){})
    group.GET("abc",function(x EngineCtx){})
    group.ANY("abc",function(x EngineCtx){})
})
```
