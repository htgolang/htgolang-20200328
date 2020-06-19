package models

import "strings"

type Task struct {
	ID int
	Name string
	Status string
}
var Taskslist = []*Task{&Task{1,"洗衣服","正在执行"},&Task{2,"做饭","未完成"}}


func AddTask(taskname string,status string)  {

	Taskslist = append(Taskslist,&Task{len(Taskslist)+1,taskname,status})
}

func Retrieve(id int) *Task {
	for _,task := range Taskslist {
		if id == task.ID{
			return task
		}
	}
	return nil
}

func Search(word string)  []*Task{
	newtask := []*Task{}
	for _,task := range Taskslist {
		if strings.Contains(task.Name,word){
			newtask = append(newtask,task)
		}
	}
	return newtask
}

func Delete(id int)  {
	for k,task := range Taskslist {
		if task.ID == id {
			Taskslist = append(Taskslist[:k],Taskslist[k+1:] ...)
		}
	}
}

func Modify(id int,name string,status string) {
	for _,task := range Taskslist {
		if task.ID == id {
			task.Name = name
			task.Status = status

		}
	}

}