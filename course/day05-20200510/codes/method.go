package main

import (
	"fmt"
	"time"
)

type Task struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	user      string
}

// name 值接收者
func (task Task) SetName(name string) {
	task.name = name
}

/* 自动生成
func (task *Task) SetName(name string) {
	(*task).SetName(name)
}
*/

// user 指针接收者
func (task *Task) SetUser(user string) {
	task.user = user
}

func main() {
	// 方法
	task := Task{}   // 值
	task2 := &Task{} //指针
	fmt.Println(task, task2)

	// 方法表达式 结构体.方法名
	// 对于值接收者 可以通过指针/值来获取方法表达式
	// GO制动针对值接收者方法 => 自动生成 指针接收者方法
	method1 := Task.SetName //func(main.Task, string)

	method1(task, "test")
	method1(*task2, "test")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	method2 := (*Task).SetName //func(*main.Task, string)

	method2(&task, "test")
	method2(task2, "test")
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	// 对于指针接收者，只能通过指针来获取方法表达式
	// method3 := Task.SetUser
	method4 := (*Task).SetUser //func(*main.Task, string)

	method4(&task, "test")
	method4(task2, "test")

	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	fmt.Printf("%T\n", method1)
	fmt.Printf("%T\n", method2)
	// fmt.Printf("%T\n", method3)
	fmt.Printf("%T\n", method4)
}
