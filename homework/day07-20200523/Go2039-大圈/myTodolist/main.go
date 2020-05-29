package main

import (
	"myTodolist/commands"
	//导入初始化包，但是在main函数中不使用init包中的方法，只是在此处调用init包中的init方法。使用 _ 导入
	_ "myTodolist/init"
)

func main() {
	//执行程序
	commands.Run()
}
/*
1. 部分输入内容未做数据验证
2. crud部分，重复代码较多。需要提高代码复用性
3. 需要多理解老师的设计思想并尝试自己实现这种思想
4. 要求接口的功能未实现
5. panic和recover的功能也暂未实现
 */