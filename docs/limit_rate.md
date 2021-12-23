# 限流器

## 令牌桶

### 配置

#### 内存

```go
os.Setenv("limiter.token", "mem://mem:10/?time_window=100") //100秒内允许10个请求
```

#### redis

```go
os.Setenv("redis.limit", "redis:xxx")
os.Setenv("limiter.token", "redis://limit:10/?time_window=100") //100秒内允许10个请求,limit 是 redis.limit的配置
```

### 使用

#### 针对接口限流

```go
func Test(ctx *web.EngineCtx) error {
ctx.LimitToken("接口路径")
}
```

#### 针对全站限流

```go
web.RegisterMiddleWare(func (ctx *web.EngineCtx, f func (*web.EngineCtx)) {//注册一个中间件
ctx.LimitToken("/app.all")
f(ctx)
})
```

## IP请求频率

### 配置

#### mem

> 无需配置

#### redis

```go
os.Setenv("redis.limitip", "redis:xxx")
os.Setenv("limiter.ip", "redis://limitip") 
```

### 使用

#### 针对接口限流

```go
func Test(ctx *web.EngineCtx) error {
ctx.LimitIP("路径", 100, time.Second * 100) //该接口100秒内一个IP可以请求100次
}
```

#### 针对全站限流

```go
web.RegisterMiddleWare(func (ctx *web.EngineCtx, f func (*web.EngineCtx)) {//注册一个中间件
ctx.LimitIP("all", 100, time.Second * 100) //100秒内一个IP可以请求100次
f(ctx)
})
```

