package models

import "fmt"

type Task struct {
	ID     int
	Name   string
	Status string
}

var tasks = []*Task{&Task{1, "洗衣服", "正在执行"}, &Task{2, "打游戏", "还没呢"}, &Task{3, "做晚餐", "马上去"}}

func GetTasks() []*Task {
	return tasks
}

func AddTask(name string) {
	tasks = append(tasks, &Task{len(tasks) + 1, name, "add..."})
}

func DelTask(id int) {
	newTasks := make([]*Task, 0)
	for _, user := range tasks {
		if user.ID == id {
			fmt.Println()
			continue
		}
		newTasks = append(newTasks, user)
	}
	tasks = newTasks
}
