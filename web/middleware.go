package web

import (
	"encoding/json"
	"github.com/ghf-go/nannan/app"
	"github.com/ghf-go/nannan/glog"
	"github.com/ghf-go/nannan/secret"
	"net/http"
	"time"
)

type BaseMiddleWare func(*EngineCtx ,func(*EngineCtx))

var (
	_middleWare = []BaseMiddleWare{JWTMiddleWare}
	_middleWareLen = 1
)

func RegisterMiddleWare(middleWare BaseMiddleWare){
	_middleWare = append(_middleWare,middleWare)
	_middleWareLen = len(_middleWare)
}
func runMiddleWare(engine *EngineCtx ,handle func(*EngineCtx)){
	if _middleWareLen == 0{
		handle(engine)
	}
	_runMiddle(engine,handle,0)

	if engine.header != nil{
		h := engine.rep.Header()
		for k,v := range engine.header{
			for _,vv := range v{
				app.Debug("end  set header %s -> %s",k,vv)
				h.Set(k,vv)
			}
		}
	}else{
		app.Debug("end not set header")
	}
	if engine.cookies != nil{
		for _,c := range engine.cookies{
			app.Debug("end： 设置 cookie %v",c)
			http.SetCookie(engine.rep,c)
		}
	}else{
		app.Debug("end： 没有设置COOKIE")
	}
	if engine.httpCode != 0{
		engine.rep.WriteHeader(engine.httpCode)
	}
	engine.rep.Write(engine.buf.Bytes())
}
func _runMiddle(engine *EngineCtx ,handle func(*EngineCtx),i int)  {
	m := _middleWare[i]
	if i + 1 == _middleWareLen{
		m(engine,handle)
	}else {
		m(engine, func(ctx *EngineCtx) {
			_runMiddle(engine,handle,i+1)
		})
	}
}
func JWTMiddleWare(engine *EngineCtx ,handle func(*EngineCtx)){
	tname := "jwt"
	aes := secret.Aes("987yhjnbgzkdlopf")
	tExpire := time.Second * 86400 * 365
	token := engine.Req.Header.Get(tname)
	engine.Session = map[string]interface{}{}
	if token == ""{
		c ,e := engine.Req.Cookie(tname)
		if e == nil{
			token = c.Value
		}else{
			app.Debug("JWT 获取COOKIE 错误 %s",e.Error())
		}
	}
	if token != ""{
		src,e := aes.Decode(token)
		if e == nil{
			data := map[string]interface{}{}
			e = json.Unmarshal([]byte(src),data)
			if e == nil{
				if ep,ok := data["expire"];ok{
					if time.Now().UnixNano() <= ep.(int64){
						delete(data,"expire")
						engine.Session = data
					}
				}
			}else{
				app.Debug("JWT JSON decode 错误 %s",e.Error())
			}
		}else{
			app.Debug("JWT Aes DECODE 错误 %s",e.Error())
		}
	}
	handle(engine)

	engine.Session["expire"] = time.Now().Add(tExpire).UnixNano()
	outJosn ,e:= json.Marshal(engine.Session)
	glog.Debug("jwt JSON ENCODE %s",string(outJosn))
	if e == nil{
		token , e := aes.Encode(string(outJosn))
		if e == nil{
			app.Debug("jwt AES ENCODE %s",token)
			engine.Header().Add(tname,token)
			engine.SetCookie(&http.Cookie{Name: tname,Value: token,Expires: time.Now().Add(tExpire),Path: "/"})
			app.Debug("jwt SET %s",token)
		}else{
			app.Debug("JWT Aes encode 错误 %s",e.Error())
		}
	}else{
		app.Debug("JWT json encode 错误 %s",e.Error())
	}
}