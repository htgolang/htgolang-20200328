package models

import "fmt"

type Task struct {
	ID     int
	Name   string
	Status string
}

var tasks = []*Task{&Task{1, "洗衣服", "正在执行"}, &Task{2, "打游戏", "还没呢"}, &Task{3, "做晚餐", "马上去"}}

//返回函数
func GetTasks() []*Task {
	return tasks
}

//增加
func AddTask(name string) {
	tasks = append(tasks, &Task{len(tasks) + 1, name, "add..."})
}

//删除
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

//修改
func EditTask(id int, name string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Name = name
		}
	}
}
