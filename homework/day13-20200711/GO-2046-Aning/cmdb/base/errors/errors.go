package errors

//定义结构体
type Errors struct {
	errors map[string][]string
}

//新增错误
func (e *Errors) Add(key, err string) {
	//判断key是否存在
	if _, ok := e.errors[key]; !ok {
		//初始化一个
		e.errors[key] = make([]string, 0, 10)
	}
	//新增到已有错误
	e.errors[key] = append(e.errors[key], err)
}

//查询所有的错误
func (e *Errors) Errors() map[string][]string {
	return e.errors
}

//只查询key对应的错误
func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}

//检查错误是否为空
func (e *Errors) HasErrors() bool {
	return len(e.errors) != 0
}

//新建一个错误
func New() *Errors {
	return &Errors{
		errors: make(map[string][]string),
	}
}
