package main

import "fmt"

type Task struct {
}

// 值接收者
func (t Task) Value() {
	fmt.Println("Value")
}

/*
func (t *Task) Value() {
	(&t)Value()
}
*/

// 指针接收者
func (t *Task) PValue() {
	fmt.Println("PValue")
}

func main() {
	var task Task
	task.Value()
	(&task).PValue()
	task.PValue() // task没有PValue方法, 语法糖

	var pTask *Task = new(Task)
	(*pTask).Value() //
	pTask.Value()    // pTask有Value方法, 语法糖
	pTask.PValue()
}
