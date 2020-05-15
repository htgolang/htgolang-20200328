package main

import (
	"fmt"
	"todolist/until/auth"
	"todolist/until/task"
)

//
//const (
//	id        = "id"
//	name      = "name"
//	startTime = "start_time"
//	endTime   = "end_time"
//	status    = "status"
//	user      = "user"
//)
//var todos = []map[string]string{
//	{"id": "1", "name": "陪孩子散步", "startTime": "18:00", "endTime": "", "status": statusNew, "user": "kk"},
//	{"id": "2", "name": "备课", "startTime": "21:00", "endTime": "", "status": statusNew, "user": "kk"},
//	{"id": "3", "name": "复习", "startTime": "09:00", "endTime": "", "status": statusNew, "user": "kk"},
//	{"id": "9", "name": "复习", "startTime": "10:00", "endTime": "", "status": statusNew, "user": "kk"},
//	{"id": "4", "name": "复习", "startTime": "10:00", "endTime": "", "status": statusNew, "user": "kk"},
//}

//const (
//	statusNew     = "未执行"
//	statusStart = "准备开始"
//	statusDoding = "进行中"
//	statusCompele = "完成"
//)

////获取todo新任务id
//func genId() int {
//	var rt int
//	for _, todo := range task.Todolist {
//		todoId := todo.Id
//		if rt < todoId {
//			rt = todoId
//		}
//	}
//	return rt + 1
//}
////显示一个任务
//func show(dataIn task.Task)  {
//	fmt.Printf("ID：%s\n",dataIn.Id)
//	fmt.Printf("NAME：%s\n",dataIn.Name)
//	fmt.Printf("STARTTIME：%s\n",dataIn.StartTime.Format(task.Layout))
//	fmt.Printf("ENDTIME：%s\n",dataIn.EndTime.Format(task.Layout))
//	fmt.Printf("STATUS：%s\n",dataIn.Status)
//	fmt.Printf("USER：%s\n",dataIn.User)
//
//}
////实例化一个新的任务
//func newTask() task.Task {
//	// id生成(用todos中最大的ID+1)
//	//task := make(map[string]string)
//	//task[id] = strconv.Itoa(genId())
//	//task[name] = ""
//	//task[startTime] = ""
//	//task[endTime] = ""
//	//task[status] = statusNew
//	//task[user] = ""
//	newtask := task.Task{
//		Id : genId(),
//		Name : "",
//		StartTime: time.Time{},
//		EndTime : time.Time{},
//		Status : "",
//		User : "",
//	}
//
//	return newtask
//}
////新增一个任务
//func add() {
//	task := newTask()
//
//	fmt.Println("请输入任务信息:")
//	for{
//		taskName := auth.Input("任务名:")
//		if !checkTaskName(taskName){
//			task.Name=taskName
//			break
//		}else {
//			fmt.Println("任务已重名，请重新输入任务名")
//		}
//	}
//
//
//	task.StartTime = time.Now()
//	task.Status = statusNew
//	task.User = auth.Input("负责人:")
//
//	task.Todolist = append(task.Todolist, task)
//	fmt.Println("创建任务成功")
//}
////编辑一个任务
//func edit()  {
//	id := auth.Input("请输入要编辑任务的id")
//	for i,j:=range todos{
//		if id == j["id"]{
//			show(j)
//			fmt.Println("请输入修改信息:")
//			for{
//				taskName := auth.Input("任务名:")
//				if !checkTaskName(taskName){
//					j[name]=taskName
//					break
//				}else {
//					fmt.Println("任务已重名，请重新输入任务名")
//				}
//			}
//			j[user] = auth.Input("负责人:")
//			{
//				for {
//					fmt.Printf("任务状态必须是一下内容之一(未执行/准备开始/进行中/完成)")
//					userInStatus := auth.Input("任务状态:")
//					if userInStatus == statusNew || userInStatus == statusStart || userInStatus == statusDoding || userInStatus == statusCompele {
//						if userInStatus == statusCompele {
//							j[status] = userInStatus
//							j[endTime] = time.Now().Format("2006-01-02 15:04:05")
//
//						}else {
//							j[status] = userInStatus
//						}
//						break
//					}else {
//						fmt.Printf("重新输入任务状态")
//					}
//				}
//			}
//			break
//		}
//	if i == len(todos)-1{
//		fmt.Println("没有这个id")
//	}
//	}
//}
////删除一个任务
//func del()  {
//	id := auth.Input("请输入要删除任务的id")
//	for i,j :=range todos{
//		if id==j["id"] {
//			show(j)
//			userInPut := auth.Input("以上是任务信息，确认要删除吗？(yes/y)")
//			switch userInPut {
//			case "yes","y":
//				if i==len(todos)-1{
//					todos=todos[0:i-1]
//				}else {
//					copy(todos[i:],todos[i+1:])
//					todos=todos[0:len(todos)-1]
//				}
//			default:
//
//			}
//			break
//		}
//		if i == len(todos)-1{
//			fmt.Println("没有这个id")
//	}
//	}
//
//}
////搜索
//func search()  {
//	id := auth.Input("请输入你要查询的id号或all")
//	if id=="all"{
//		for _,j:=range todos{
//			fmt.Println(strings.Repeat("-",20))
//			show(j)
//			fmt.Println(strings.Repeat("-",20))
//		}
//	}else {
//		for i,j:=range todos{
//			if id==j["id"]{
//				show(j)
//				break
//			}
//			if i==len(todos)-1{
//				fmt.Println("没有这个id")
//			}
//		}
//	}
//
//}
////退出
//func exit()  {
//	os.Exit(0)
//}
////验证任务名
//func checkTaskName(name string) bool  {
//	for i,j :=range todos{
//		if j["name"]==name{
//			return true
//		}
//		if i==len(todos)-1 {
//			return false
//		}
//		}
//	return false
//	}
//
//
//func showall()  {
//	sortrules:="id"
//EXIT:	for {
//		fmt.Printf("以%s排序\n",sortrules)
//		autoFmt:=false
//		autoWrp:=false
//		reflows:=false
//		t := tablewriter.NewWriter(os.Stdout)
//		t.SetAutoFormatHeaders(autoFmt)
//		t.SetAutoWrapText(autoWrp)
//		t.SetReflowDuringAutoWrap(reflows)
//		title := []string{"id","name","startTime","endTime","status","user"}
//		if sortrules=="id"{		sort.SliceStable(todos, func(i, j int) bool {
//			a,_:=strconv.Atoi(todos[i]["id"])
//			b,_:=strconv.Atoi(todos[j]["id"])
//			return a<b
//
//		})}else {
//			sort.SliceStable(todos, func(i, j int) bool {
//
//				return todos[i][sortrules]<todos[j][sortrules]
//			})
//		}
//		t.SetHeader(title)
//		for _,j:=range todos{
//			t.Append([]string{j[title[0]],j[title[1]],j[title[2]],j[title[3]],j[title[4]],j[title[5]],})
//		}
//		t.Render()
//		for {
//			have:=false
//			fmt.Println("请输入排序规则(id/name/exit/startTime)")
//			fmt.Scan(&sortrules)
//			if sortrules=="exit"{
//				break EXIT
//			}
//			for _,j:=range title{
//				if j ==sortrules{
//					have =true
//				}
//			}
//			if !have{
//				fmt.Println("没有这个字段，无法排序")
//			}else {
//				break
//			}
//
//		}
//
//	}
//
//
//
//}
func main() {


	auth.Auth()
	methods := map[string]func(){
		"add":task.Add,
		"edit":task.Edit,
		"del":task.Del,
		"search":task.Search,
		"exit":task.Exit,
		"showall":task.Showall,
	}
	for {
		userinput := auth.Input("请输入你的操作(add/edit/del/search/exit/showall)")
		if method,ok:=methods[userinput];ok{
			method()
		}else {
			fmt.Println("没有该指令，请重新输入")
		}
	}


}