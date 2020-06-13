package models

type Task struct {
	ID     int
	Name   string
	Status string
}

var tasks = []*Task{&Task{1, "洗衣服", "正在执行"}, &Task{2, "做作业", "未执行"}}

func GetTasks() []*Task {
	return tasks
}

func AddTask(name string) {
	tasks = append(tasks, &Task{len(tasks), name, "新增"})
}
