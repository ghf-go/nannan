# 中间件

## 注册中间件

```go
web.RegisterMiddleWare(func (ctx *web.EngineCtx, f func (*web.EngineCtx)) {
//before处理
gerr.Error(500, "服务器繁忙") //报错退出功能
f(ctx)
//after处理
})
```

## JWT

```go
web.RegisterMiddleWare(web.JWTMiddleWare)
```

## 微信验证

```go
web.RegisterMiddleWare(web.WxEchoStrMiddkeWare)
```

## Session