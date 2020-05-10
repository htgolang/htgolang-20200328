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

func NewTask(id int, name string, user string) *Task {
	return &Task{
		id:        id,
		name:      name,
		startTime: time.Now(),
		user:      user,
	}
}

// name 值接收者
func (task Task) SetName(name string) {
	task.name = name
}

// user 指针接收者
func (task *Task) SetUser(user string) {
	task.user = user
}

func main() {
	task := Task{name: "完成TODO"}

	task.SetName("知识整理") // 方法调用

	// 语法糖 go编译时 取引用
	task.SetUser("kk") // 可以 (&task).SetUser("kk")

	fmt.Printf("%#v\n", task)

	task2 := &Task{name: "完成TODO"}

	// 语法糖(虽然是通过指针调用，但并会修改name执行)
	// 解引用
	task2.SetName("知识整理") // (*task2).SetName("知识整理")
	task2.SetUser("kk")

	fmt.Printf("%#v\n", task2)

	var task3 Task  // 值
	var task4 *Task // 指针
	task3.SetName("")
	task4.SetName("")
}
