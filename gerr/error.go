package gerr

import "reflect"

type BaseErr struct {
	Code int
	Msg  string
}

func Error(code int, msg string) {
	panic(&BaseErr{
		Code: code,
		Msg:  msg,
	})
}
func IsError(v interface{}) bool {
	t := reflect.TypeOf(v)
	return t.String() == "gerr._error"
}
func (r *BaseErr) Error() string {
	return r.Msg
}
