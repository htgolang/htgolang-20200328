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

type Task struct {
	id        int
	name      string
	startTime time.Time
	endTime   time.Time
	status    int
	*User
}

func main() {
	var task Task

	fmt.Printf("%#v\n", task)
	task = Task{
		id:   1,
		name: "完成todo",
		User: &User{
			id:   1,
			name: "kk",
			addr: "西安",
		},
	}
	fmt.Printf("%#v\n", task)
	fmt.Println(
		task.name,
		task.User.name,
		task.User.addr,
		task.addr,
	)
}
