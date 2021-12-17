package gutils

import "reflect"

type _error struct {
	Code int
	Msg  string
}

func Error(code int, msg string) {
	panic(_error{
		Code: code,
		Msg:  msg,
	})
}

func IsError(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.String() == "gutils._error"
}
