package web

import (
	"encoding/json"
	"github.com/ghf-go/nannan/glog"
	"github.com/ghf-go/nannan/secret"
	"net/http"
	"time"
)

type BaseMiddleWare func(*EngineCtx, func(*EngineCtx))

var (
	_middleWare    = []BaseMiddleWare{}
	_middleWareLen = 0
)

func RegisterMiddleWare(middleWare BaseMiddleWare) {
	_middleWare = append(_middleWare, middleWare)
	_middleWareLen = len(_middleWare)
}
func runMiddleWare(engine *EngineCtx, handle func(*EngineCtx)) {
	if _middleWareLen == 0 {
		handle(engine)
	}
	_runMiddle(engine, handle, 0)

	if engine.header != nil {
		h := engine.rep.Header()
		for k, v := range engine.header {
			for _, vv := range v {
				h.Set(k, vv)
			}
		}
	}
	if engine.cookies != nil {
		for _, c := range engine.cookies {
			http.SetCookie(engine.rep, c)
		}
	}
	if engine.httpCode != 0 {
		engine.rep.WriteHeader(engine.httpCode)
	}
	engine.rep.Write(engine.buf.Bytes())
}
func _runMiddle(engine *EngineCtx, handle func(*EngineCtx), i int) {
	m := _middleWare[i]
	if i+1 == _middleWareLen {
		m(engine, handle)
	} else {
		m(engine, func(ctx *EngineCtx) {
			_runMiddle(engine, handle, i+1)
		})
	}
}
func JWTMiddleWare(engine *EngineCtx, handle func(*EngineCtx)) {
	tname := "jwt"
	aes := secret.Aes("987yhjnbgzkdlopf")
	tExpire := time.Second * 86400 * 365
	token := engine.Req.Header.Get(tname)
	engine.session = session{}
	if token == "" {
		c, e := engine.Req.Cookie(tname)
		if e == nil {
			token = c.Value
		} else {
			//glog.AppDebug("JWT 获取COOKIE 错误 %s",e.Error())
		}
	}
	if token != "" {
		src, e := aes.Decode(token)
		if e == nil {
			d1 := make(session)
			e = json.Unmarshal([]byte(src), &d1)
			if e == nil {
				if ep, ok := d1["expire"]; ok {
					if time.Now().Unix() <= int64(ep.(float64)) {
						delete(d1, "expire")
						engine.session = d1
					}

				}
			} else {
				glog.AppDebug("JWT JSON decode 错误 %s -> (%s)", e.Error(), src)
			}
		} else {
			glog.AppDebug("JWT Aes DECODE 错误 %s", e.Error())
		}
	}
	handle(engine)
	engine.SetSession("expire", time.Now().Add(tExpire).Unix())
	outJosn, e := json.Marshal(engine.session)
	if e == nil {
		token, e := aes.Encode(string(outJosn))
		if e == nil {
			engine.Header().Add(tname, token)
			engine.SetCookie(&http.Cookie{Name: tname, Value: token, Expires: time.Now().Add(tExpire), Path: "/"})
		} else {
			glog.AppDebug("JWT Aes encode 错误 %s", e.Error())
		}
	} else {
		glog.AppDebug("JWT json encode 错误 %s", e.Error())
	}
}

//微信服务器校验
func WxEchoStrMiddkeWare(engine *EngineCtx, handle func(*EngineCtx)) {
	ecstr := engine.get("echostr")
	if ecstr != "" {
		engine.Write([]byte(ecstr))
		return
	}
	handle(engine)
}
func RedisSssionMiddkeWare(engine *EngineCtx, handle func(*EngineCtx)) {

}
