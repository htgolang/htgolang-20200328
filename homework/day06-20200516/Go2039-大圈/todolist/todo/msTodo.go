package todo

//初始化一个字典
var msTodoMap = make(map[string][]*todo)

//创建一个管理todolist的结构体
type msTodo struct {
	todoItems map[string][]*todo
}

func NewmsT() *msTodo {
	//实例化msTodo结构体
	return &msTodo{
		todoItems: msTodoMap,
	}
}
//实例化一个任务管理系统实例
var msgTodo = NewmsT()