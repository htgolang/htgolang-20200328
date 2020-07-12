package baseerr

import "github.com/astaxie/beego"

type Errors struct {
	errors map[string][]string
}

func (e *Errors) Add(key , err string) {
	if _,ok := e.errors[key];!ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key],err)
	beego.Error(e.errors)
}
//返回错误列表
func (e *Errors) Errors() map[string][]string {
	return e.errors
}

func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}
//如果他的切片长度不等于0 那么就返回true  说明有错误
//反之没错
func (e *Errors) HasErrors() bool {
	return  len(e.errors) != 0
}

func New() *Errors {
	return &Errors{
		errors: make(map[string][]string),
	}
}