# nannan

# 使用

## CLI

> 注册
>
> app.RegisterCli(name string, callfunc func([]string))
>
> 运行 main cli cliname

## Crontanb

> 注册
> app.RegisterCrontan(name, timer string, isLock bool, callfunc func())
> 运行 main crontab  name [save|remove|check]
>
> save 保存计划任务
>
> remove 删除计划任务
>
> check 查看计划任务

## Service

> 注册
> RegisterService(name string, callfunc func())
>
> 编译程序后执行 main service servicename [install|uninstall|start|stop|restart]
> install 安装服务
>
> uninstall 删除服务
>
> start 启动服务
>
> stop 关闭服务

## Web

> 程序编译之后直接运行即可，默认是web应用

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
web.RegisterRouterGroup("/admin", func (group *RouterGroup){
group.POST("abc", function(x EngineCtx){})
group.GET("abc", function(x EngineCtx){})
group.ANY("abc", function(x EngineCtx){})
})
```
