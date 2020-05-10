package main

import (
	"fmt"
	"time"
)

type User struct {
	id   int
	name string
	addr string
	tel  string
}

type Creator struct {
	id   int
	name string
	addr string
}

type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	User      // 嵌入(匿名) => 面向对象(继承) 也是有一个属性名字 User(默认简写)
	Creator
}

func main() {
	var task Task
	fmt.Printf("%#v\n", task)

	fmt.Println(task.User.name)
	task.User.name = "kk"
	fmt.Printf("%#v\n", task)

	// 赋值
	task = Task{
		id:   1,
		name: "完成TODOLISt",
		User: User{
			id:   1,
			name: "大全",
			addr: "北京",
		},
	}
	fmt.Printf("%#v\n", task)

	fmt.Println(task.name)
	fmt.Println(task.Creator.addr)

	task.Creator.addr = "新加坡"
	fmt.Println(task.Creator.addr)
	fmt.Printf("%#v\n", task)

	task.User.addr = "新加坡2"
	fmt.Println(task.User.addr)
	fmt.Printf("%#v\n", task)

}
