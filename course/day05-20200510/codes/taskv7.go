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

func NewUser(id int, name, addr, tel string) *User {
	return &User{id, name, addr, tel}
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetAddr(addr string) {
	u.addr = addr
}

type Task struct {
	id        int
	name      string
	startTime *time.Time
	endTime   *time.Time
	status    int
	*User
}

func NewTask(id int, name string, startTime, endTime *time.Time, user *User) *Task {
	return &Task{id, name, startTime, endTime, 1, user}
}

func main() {
	start := time.Now()
	end := start.Add(24 * time.Hour)

	user := NewUser(1, "kk", "", "")
	task := NewTask(1, "完成知识整理", &start, &end, user)

	task.SetAddr("北京") // task.User.SetAddr
	task.SetName("大圈") // task.User.SetName

	fmt.Printf("%#v\n", task)
	fmt.Printf("%#v\n", task.User)
}
