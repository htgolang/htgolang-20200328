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

// 定义方法
// 特定类型指定 => 接收者
func (task *Task) SetName(name string) {
	task.name = name
}

func (task *Task) GetName() string {
	return task.name
}

func main() {
	task := &Task{name: "完成TODO"}
	task.SetName("知识整理") // 方法调用

	fmt.Println(task.GetName()) // 知识整理
}
