package main

import "newtodolist/todo"

func main() {
	todo.ChoiceNew()
}

/*
bug1: 要求任务ID唯一，但是从文件里读取数据时并为做验证ID唯一性
bug2: gob格式存储数据时正常，但是反序列化时只能读取第一行数据
bug3: 输入all的时候，需要打印出所有的信息。但此功能暂未写
 */