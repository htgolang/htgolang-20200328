package operations

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"todolist/task"
)
const (
	statusNew = 0
	statusComplete = 1
	statusIncomplete = 2
)
var user1 = task.NewUser("chen","Singapore","123456")
var user2 = task.NewUser("Jason","Beijing","98765")
var start = time.Now()
//var end = start.Add(-24*time.Hour)
var  Todols= []*task.Task{
	task.NewTask(1,"读书",&start,statusNew,user1),
	task.NewTask(2,"洗澡",&start,statusNew,user2),
	task.NewTask(3,"读报纸",&start,statusNew,user1),
	task.NewTask(4,"上网",&start,statusNew,user2),
	task.NewTask(5,"上山",&start,statusNew,user1),
	task.NewTask(6,"洗衣服",&start,statusNew,user1),
	task.NewTask(7,"上厕所",&start,statusNew,user1),
	task.NewTask(8,"洗电脑",&start,statusNew,user2),
	task.NewTask(9,"上超市",&start,statusNew,user1),
	task.NewTask(10,"读圣经",&start,statusNew,user2),
}
func genId(tasks []*task.Task) int  {
	var max int
	for _,todo := range tasks{
		if todo.GetIdTask() > max{
			max = todo.GetIdTask()
		}
	}
	return max + 1
}
func Input(prompt string) string  {
	var text string
	fmt.Println(prompt)
	fmt.Scan(&text)
	return text
}
func Add(tasks []*task.Task)   []*task.Task {
	new_id := genId(tasks)
	new_name := Input("请输入任务名字:")
	new_startTime := time.Now()
	new_pstartTime := &new_startTime
	new_status := statusNew
	new_username := Input("请输入用户名:")
	new_addr := Input("请输入地址:")
	new_tel := Input("请输入电话号码:")
	usr := task.NewUser(new_username,new_addr,new_tel)
	todo := task.NewTask(new_id,new_name,new_pstartTime,new_status,usr)
	 tasks = append(tasks,todo)
	fmt.Println("添加成功！")
	fmt.Println(todo,*todo.User)
	 return tasks

}
func Query(tasks []*task.Task)  []*task.Task  {
	var ls []*task.Task
	char := Input("请输入查询关键字")
	for _,todo := range tasks{
		if strings.Contains(todo.GetNameTask(),char){
			ls = append(ls,todo)
		}
	}
	fmt.Println(ls)
	return ls
}
func Modify(tasks []*task.Task)  []*task.Task {
	var query_id = Input("请输入任务的ID")
	var flag bool
	id,_ := strconv.Atoi(query_id)
	for _,todo :=range tasks{
		if id == todo.GetIdTask(){
			flag = true
			fmt.Println(todo,*todo.User)
			var text = Input("确认要修改吗?(y/n)")
			if text == "y"{
				var content = Input("请输入修改信息:")
				todo.SetNameTask(content)
				var new_status = Input("请输入任务状态(新任务:0/完成:1/未完成:2) :")
				stat,_ := strconv.Atoi(new_status)
				todo.SetStatusTask(stat)
				if stat ==1 {
					end := time.Now()
					todo.SetendTime(&end)
				}
				fmt.Println("修改成功！")
				fmt.Println(todo,*todo.User)
			} else {
				return tasks
			}
		}
	}
	if !flag{
		fmt.Println("此ID不存在！")
	}
	return tasks
}

func Delete(tasks []*task.Task) []*task.Task {
	var delete_id = Input("请输入要删除的任务ID:")
	id,_ :=strconv.Atoi(delete_id)
	for idx,todo := range tasks{
		if id == todo.GetIdTask() {
			fmt.Println(todo,*todo.User)
			var text = Input("确认删除吗？(y/n)")
			if text == "y" {
				copy(tasks[idx:],tasks[idx+1:])
				tasks = Todols[:len(tasks)-1]
				fmt.Println("删除成功！")
				fmt.Println(tasks)
			}else{
				return tasks
			}
		}
	}
	return tasks
}
