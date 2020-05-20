package main
import (
	"fmt"
	"todolist/operations"
	"todolist/task"
	"todolist/passwordfile"
	"todolist/saving"
	"todolist/reading"
)

//var end = start.Add(-24*time.Hour)
//var tasks = operations.Todols

func basic_op(tasks []*task.Task)  []*task.Task{
	methods := map[string]func(tasks []*task.Task) []*task.Task{
		"add":operations.Add,
		"query":operations.Query,
		"modify":operations.Modify ,
		"delete":operations.Delete ,
	}
	for {
		text := operations.Input("请输入操作(add/query/modify/delete/exit):")
		if text == "exit" {
			return tasks
		}
		if method, ok := methods[text]; ok {
			tasks =method(tasks)
		} else {
			fmt.Println("输入指令不正确！")
		}
	}
	return tasks
}
func main()  {
	pass_file := "password.txt"
	var pass string
	var final_tasks []*task.Task

	tasks := reading.Reading_option()
	fmt.Printf("%#v\n",tasks)
	pass = operations.Input("请输入密码:")
	if passwordfile.FileIsExists(pass_file){
		existing_pass := passwordfile.ReadPwFile(pass_file)
		if passwordfile.Check_password_nosalt(existing_pass,pass){
			final_tasks = basic_op(tasks)
			fmt.Println(final_tasks)
		}else{
			fmt.Println("密码不正确！")
		}
	}else{
		passwordfile.WritePwFile(pass_file,pass)
		final_tasks = basic_op(tasks)
	}

	fmt.Printf("%#v\n",final_tasks)
	saving.Saving_option(final_tasks)
}
