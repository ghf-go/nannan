package web

import (
	"encoding/json"
	"fmt"
	"github.com/ghf-go/nannan/gerr"
	"github.com/ghf-go/nannan/verify"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type EngineCtx struct {
	*gresponse
	session
	ReqID     int64
	Req       *http.Request
	rep       http.ResponseWriter
	ip        string
	GroupPath string
	NodePath  string
	_get_data url.Values
}

// 输出JSON
func (engine *EngineCtx) json(code int, msg string, data interface{}) error {
	ret := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	b, e := json.Marshal(ret)
	if e == nil {
		engine.Header().Set("Content-Type", "application/json; charset=utf-8")
		engine.Write(b)
	}
	return e
}

//输出错误的json
func (engine *EngineCtx) JsonFail(code int, msg string) error {
	return engine.json(code, msg, nil)
}

//输出正确的json
func (engine *EngineCtx) JsonSuccess(data interface{}) error {
	return engine.json(0, "", data)
}

//显示网页
func (engine *EngineCtx) Display(tpl string, data interface{}) error {
	return nil
}

// 获取IP
func (engine *EngineCtx) GetIP() string {
	if engine.ip != "" {
		return engine.ip
	}
	xForwardedFor := engine.Req.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(engine.Req.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(engine.Req.RemoteAddr)); err == nil {
		return ip
	}

	return "0.0.0.0"
}

func (engine *EngineCtx) Verify(obj interface{}) {
	errCode := 401
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		gerr.Error(errCode, "请求参数必须是引用类型")
	}
	if t.Elem().Kind() != reflect.Struct {
		gerr.Error(errCode, "请求参数必须是结构体")
	}
	t = t.Elem()
	fl := t.NumField()
	vo := reflect.ValueOf(obj).Elem()
	for i := 0; i < fl; i++ {
		f := t.Field(i)
		tag := f.Tag
		rk := tag.Get("get")
		sv := ""
		if rk != "" {
			sv = engine.get(rk)
		} else {
			rk = tag.Get("post")
			if rk != "" {
				sv = engine.Req.FormValue(rk)
			}
		}
		o := vo.Field(i)
		strSaveVal(sv, f, o)
		tv := tag.Get("verify")
		tvs := strings.Split(tv, ";")
		for _, vn := range tvs {
			switch vn {
			case "required":
				if sv == "" {
					gerr.Error(errCode, rk+"参数必填")
				}
			case "email":
				if !verify.IsEmail(sv) {
					gerr.Error(errCode, rk+"参数必须是邮箱地址")
				}
			case "mobile":
				if !verify.IsMobile(sv) {
					gerr.Error(errCode, rk+"参数必须是手机号")
				}
			case "date":
				if !verify.IsDate(sv) {
					gerr.Error(errCode, rk+"参数必须是日期")
				}
			case "time":
				if !verify.IsTime(sv) {
					gerr.Error(errCode, rk+"参数必须是时间")
				}
			case "datetime":
				if !verify.IsDateTime(sv) {
					gerr.Error(errCode, rk+"参数必须是日期时间")
				}
			case "url":
				if !verify.IsUrl(sv) {
					gerr.Error(errCode, rk+"参数必须是url")
				}
			case "ipv4":
				if !verify.IsIPv4(sv) {
					gerr.Error(errCode, rk+"参数必须是IP地址")
				}
			case "ipv6":
				if !verify.IsIPv6(sv) {
					gerr.Error(errCode, rk+"参数必须是IP地址")
				}
			default:
				if strings.HasPrefix(vn, "max:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[4:], 10, 64)
						if max < o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[4:], 10, 64)
						if max < o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[4:], 64)
						if max < o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%f", rk, max))
						}
					}
				} else if strings.HasPrefix(vn, "min:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[4:], 10, 64)
						if max > o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[4:], 10, 64)
						if max > o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[4:], 64)
						if max > o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%f", rk, max))
						}
					}

				} else if strings.HasPrefix(vn, "in:") {
					arr := strings.Split(vn[3:], ",")
					isOk := false
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						for _, v1 := range arr {
							v2, _ := strconv.ParseInt(v1, 10, 64)
							if v2 == o.Int() {
								isOk = true
							}
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						for _, v1 := range arr {
							v2, _ := strconv.ParseUint(v1, 10, 64)
							if v2 == o.Uint() {
								isOk = true
							}
						}
					case reflect.Float32, reflect.Float64:
						for _, v1 := range arr {
							v2, _ := strconv.ParseFloat(v1, 64)
							if v2 == o.Float() {
								isOk = true
							}
						}
					case reflect.String:
						for _, v1 := range arr {
							if v1 == o.String() {
								isOk = true
							}
						}
					}
					if !isOk {
						gerr.Error(errCode, fmt.Sprintf("%s 必须在 %s", rk, vn[3:]))
					}
				} else if strings.HasPrefix(vn, "gt:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max >= o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max >= o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max >= o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于%f", rk, max))
						}
					}
				} else if strings.HasPrefix(vn, "ge:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max > o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max > o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max > o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须大于等于%f", rk, max))
						}
					}
				} else if strings.HasPrefix(vn, "lt:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max <= o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max <= o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max <= o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于%f", rk, max))
						}
					}
				} else if strings.HasPrefix(vn, "le:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max < o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max < o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max < o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须小于等于%f", rk, max))
						}
					}
				} else if strings.HasPrefix(vn, "eq:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max != o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max != o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max != o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须等于%f", rk, max))
						}
					case reflect.String:
						if vn[3:] != o.String() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须等于%s", rk, vn[3:]))
						}
					}

				} else if strings.HasPrefix(vn, "ne:") {
					switch f.Type.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						max, _ := strconv.ParseInt(vn[3:], 10, 64)
						if max == o.Int() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须不等于%d", rk, max))
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						max, _ := strconv.ParseUint(vn[3:], 10, 64)
						if max == o.Uint() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须不等于%d", rk, max))
						}
					case reflect.Float32, reflect.Float64:
						max, _ := strconv.ParseFloat(vn[3:], 64)
						if max == o.Float() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须不等于%f", rk, max))
						}
					case reflect.String:
						if vn[3:] == o.String() {
							gerr.Error(errCode, fmt.Sprintf("%s 必须不等于%s", rk, vn[3:]))
						}
					}
				}
			}
		}

	}
}

func (engine *EngineCtx) get(key string) string {
	if engine._get_data == nil {
		engine._get_data = engine.Req.URL.Query()
	}
	return engine._get_data.Get(key)
}

func strSaveVal(src string, t reflect.StructField, o reflect.Value) error {
	switch t.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, e := strconv.ParseInt(src, 10, 64)
		if e != nil {
			return e
		}
		o.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, e := strconv.ParseUint(src, 10, 64)
		if e != nil {
			return e
		}
		o.SetUint(i)

	case reflect.Float32, reflect.Float64:
		i, e := strconv.ParseFloat(src, 64)
		if e != nil {
			return e
		}
		o.SetFloat(i)
	case reflect.Bool:
		i, e := strconv.ParseBool(src)
		if e != nil {
			return e
		}
		o.SetBool(i)
	case reflect.String:
		o.SetString(src)
	}
	return nil
}
