package main

import (
	"fmt"
	"time"
)

// 定义Task结构体类型
type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	user      string
}

func main() {
	// var name type
	var task Task // 结构体类型的变量=> 结构体的实例

	fmt.Printf("%T\n", task)
	fmt.Printf("%#v\n", task)

	// 赋值
	// 零值
	task = Task{}
	fmt.Printf("%#v\n", task)

	// 按照属性顺序创建Task字面量进行赋值
	task = Task{1, "完成Todolist", time.Now(), time.Now().Add(24 * time.Hour), 1, "kk"}
	fmt.Printf("%#v\n", task)

	// 按照属性名进行初始化
	task = Task{
		id:   2,
		name: "完成知识整理任务",
		user: "kk", // 注意最后一个元素后的,不能省略
	}

	fmt.Printf("%#v\n", task)

	// nums := map[string]int{
	// 	"1": 1, "2": 2, "3": 3,
	// }
	// fmt.Println(nums)
	// nums2 := []int{
	// 	1, 2, 3,
	// }
	// fmt.Println(nums2)

	var ptask *Task
	fmt.Printf("%T\n", ptask)
	fmt.Printf("%#v\n", ptask)

	// 赋值
	ptask = &Task{}
	fmt.Printf("%#v\n", ptask)

	// 属性的访问/修改
	fmt.Println(task.id, task.status, task.name, task.user)

	task.status = 3
	fmt.Printf("%#v\n", task)

	// 值类型
	task2 := task

	task2.user = "大圈"
	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task2)

	change(task)
	fmt.Printf("%#v\n", task) //? task.user 变不变？

	// 初始化结构体(指针)方式
	ptask2 := new(Task)
	fmt.Printf("%T\n", ptask2)
	fmt.Printf("%#v\n", ptask2)
}

func change(task Task) {
	task.user = "xxxxxxxxxxxxxxx"
}
