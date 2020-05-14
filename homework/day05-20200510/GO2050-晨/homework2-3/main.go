package main
import (
	"fmt"
	"todolist/operations"
	//"todolist/task"
	"todolist/passwordfile"
)

//var end = start.Add(-24*time.Hour)
func basic_op()  {
	methods := map[string]func(){
		"add":operations.Add,
		"query":operations.Query,
		"modify":operations.Modify,
		"delete":operations.Delete,
	}
	for {
		text := operations.Input("请输入操作(add/query/modify/delete/exit):")
		if text == "exit" {
			return
		}
		if method, ok := methods[text]; ok {
			method()
		} else {
			fmt.Println("输入指令不正确！")
		}
	}
}
func main()  {
	pass_file := "password.txt"
	var pass string
	pass = operations.Input("请输入密码:")
	if passwordfile.FileIsExists(pass_file){
		existing_pass := passwordfile.ReadPwFile(pass_file)
		if passwordfile.Check_password_nosalt(existing_pass,pass){
			basic_op()
		}else{
			fmt.Println("密码不正确！")
		}
	}else{
		passwordfile.WritePwFile(pass_file,pass)
		basic_op()
	}

}
