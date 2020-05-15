package task

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"todolist/until/auth"
)

const (
	StatusNew     = "未执行"
	StatusStart = "准备开始"
	StatusDoding = "进行中"
	StatusCompele = "完成"
)

const Layout string = "2006-01-02 15:04:05"

var Todolist []Task

type Task struct {
	Id int
	Name string
	StartTime time.Time
	EndTime time.Time
	Status string
	User string
}

func init()  {
	LoadTask()
}

func LoadTask()  {
	file,_ := os.OpenFile("task.txt",os.O_RDONLY,os.ModePerm)
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan(){
		tmp :=scaner.Text()
		tmpSlice := strings.Split(tmp,",")
		id,_ := strconv.Atoi(tmpSlice[0])
		name :=tmpSlice[1]
		startTime,_ := time.Parse(Layout,tmpSlice[2])
		endTime,_ := time.Parse(Layout,tmpSlice[3])
		status := tmpSlice[4]
		user := tmpSlice[5]
		Todolist =append(Todolist,Task{
			Id: id,
			Name:name,
			StartTime:startTime,
			EndTime:endTime,
			Status:status,
			User:user,
		})
	}

}

func WriteTask()  {
	file,_ := os.OpenFile("task.txt",os.O_RDONLY|os.O_TRUNC|os.O_CREATE,os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _,task :=range Todolist{
		id := strconv.Itoa(task.Id)
		name := task.Name
		startTime := task.StartTime.Format(Layout)
		endTime := task.EndTime.Format(Layout)
		status := task.Status
		user := task.User
		strLin := strings.Join([]string{id,name,startTime,endTime,status,user},",")
		writer.WriteString(strLin+"\n")
		writer.Flush()
	}

}



func genId() int {
	var rt int
	for _, todo := range Todolist {
		todoId := todo.Id
		if rt < todoId {
			rt = todoId
		}
	}
	return rt + 1
}
//显示一个任务
func show(dataIn Task)  {
	fmt.Printf("ID：%d\n",dataIn.Id)
	fmt.Printf("NAME：%s\n",dataIn.Name)
	fmt.Printf("STARTTIME：%s\n",dataIn.StartTime.Format(Layout))
	fmt.Printf("ENDTIME：%s\n",dataIn.EndTime.Format(Layout))
	fmt.Printf("STATUS：%s\n",dataIn.Status)
	fmt.Printf("USER：%s\n",dataIn.User)

}
//实例化一个新的任务
func newTask() Task {
	// id生成(用todos中最大的ID+1)
	//task := make(map[string]string)
	//task[id] = strconv.Itoa(genId())
	//task[name] = ""
	//task[startTime] = ""
	//task[endTime] = ""
	//task[status] = statusNew
	//task[user] = ""
	newtask := Task{
		Id : genId(),
		Name : "",
		StartTime: time.Time{},
		EndTime : time.Time{},
		Status : "",
		User : "",
	}

	return newtask
}
//新增一个任务
func Add() {
	task := newTask()

	fmt.Println("请输入任务信息:")
	for{
		taskName := auth.Input("任务名:")
		if !checkTaskName(taskName){
			task.Name=taskName
			break
		}else {
			fmt.Println("任务已重名，请重新输入任务名")
		}
	}


	task.StartTime = time.Now()
	task.Status = StatusNew
	task.User = auth.Input("负责人:")

	Todolist = append(Todolist, task)
	fmt.Println("创建任务成功")
	defer WriteTask()
}

func Edit()  {
	defer WriteTask()
	id,err := strconv.Atoi(auth.Input("请输入要编辑任务的id"))
	if err != nil{
		fmt.Println("输入有误")
		return
	}
	for i,j:=range Todolist{
		if id == j.Id{
			show(j)
			fmt.Println("请输入修改信息:")
			for{
				taskName := auth.Input("任务名:")
				if !checkTaskName(taskName){
					j.Name=taskName
					break
				}else if j.Name == taskName{
					break
				}else{
					fmt.Println("任务已重名，请重新输入任务名")
				}
			}
			j.User = auth.Input("负责人:")
			{
				for {
					fmt.Printf("任务状态必须是一下内容之一(未执行/准备开始/进行中/完成)")
					userInStatus := auth.Input("任务状态:")
					if userInStatus == StatusNew || userInStatus == StatusStart || userInStatus == StatusDoding || userInStatus == StatusCompele {
						if userInStatus == StatusCompele {
							j.Status = userInStatus
							j.EndTime = time.Now()
							fmt.Println(userInStatus,StatusCompele)
							fmt.Println(1)

						}else {
							j.Status = userInStatus
							fmt.Println(2)
						}
						break
					}else {
						fmt.Printf("重新输入任务状态")
					}
				}
			}
			break
		}
		if i == len(Todolist)-1{
			fmt.Println("没有这个id")
		}
		defer func() {Todolist[i]=j}()
	}

}
//删除一个任务
func Del()  {
	defer WriteTask()
	id,err := strconv.Atoi(auth.Input("请输入要删除任务的id"))
	if err !=nil{
		fmt.Println("输入有误")
		return
	}
	for i,j :=range Todolist{
		if id==j.Id {
			show(j)
			userInPut := auth.Input("以上是任务信息，确认要删除吗？(yes/y)")
			switch userInPut {
			case "yes","y":
				if i==len(Todolist)-1{
					Todolist=Todolist[0:i-1]
				}else {
					copy(Todolist[i:],Todolist[i+1:])
					Todolist=Todolist[0:len(Todolist)-1]
				}
			default:

			}
			break
		}
		if i == len(Todolist)-1{
			fmt.Println("没有这个id")
		}
		defer func() {Todolist[i]=j}()
	}

}
//搜索
func Search()  {
	id,err := strconv.Atoi(auth.Input("请输入你要查询的id号或-1(全部)"))
	if err !=nil{
		fmt.Println("输入有误")
		return
	}
	if id<0{
		for _,j:=range Todolist{
			fmt.Println(strings.Repeat("-",20))
			show(j)
			fmt.Println(strings.Repeat("-",20))
		}
	}else {
		for i,j:=range Todolist{
			if id==j.Id{
				show(j)
				break
			}
			if i==len(Todolist)-1{
				fmt.Println("没有这个id")
			}
		}
	}

}
//退出
func Exit()  {
	os.Exit(0)
}
//验证任务名
func checkTaskName(name string) bool  {
	for i,j :=range Todolist{
		if j.Name==name{
			return true
		}
		if i==len(Todolist)-1 {
			return false
		}
	}
	return false
}


func Showall()  {
	sortrules:="id"
EXIT:	for {
	fmt.Printf("以%s排序\n",sortrules)
	autoFmt:=false
	autoWrp:=false
	reflows:=false
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(autoFmt)
	t.SetAutoWrapText(autoWrp)
	t.SetReflowDuringAutoWrap(reflows)
	title := []string{"id","name","startTime","endTime","status","user"}
	if sortrules=="id"{		sort.SliceStable(Todolist, func(i, j int) bool {
		a:=Todolist[i].Id
		b:=Todolist[j].Id
		return a<b

	})}else if sortrules == "name" {
		sort.SliceStable(Todolist, func(i, j int) bool {
			return Todolist[i].Name<Todolist[j].Name
		})
	}else if sortrules == "startTime"{
		sort.SliceStable(Todolist, func(i, j int) bool{
			return Todolist[i].StartTime.After(Todolist[j].StartTime)
		})
	}else if sortrules == "endTime"{
		sort.SliceStable(Todolist, func(i, j int) bool{
			return Todolist[i].EndTime.After(Todolist[j].EndTime)
		})
	}else if sortrules == "status" {
		sort.SliceStable(Todolist, func(i, j int) bool {
			return Todolist[i].Status<Todolist[j].Status
		})
	}else if sortrules == "user" {
		sort.SliceStable(Todolist, func(i, j int) bool {
			return Todolist[i].User<Todolist[j].User
		})
	}
	t.SetHeader(title)
	for _,j:=range Todolist{
		id := strconv.Itoa(j.Id)
		name := j.Name
		startTime := j.StartTime.Format(Layout)
		endTime := j.EndTime.Format(Layout)
		status := j.Status
		user := j.User
		t.Append([]string{id,name,startTime,endTime,status,user})
	}
	t.Render()
	for {
		have:=false
		fmt.Println("请输入排序规则(id/name/exit/startTime)")
		fmt.Scan(&sortrules)
		if sortrules=="exit"{
			break EXIT
		}
		for _,j:=range title{
			if j ==sortrules{
				have =true
			}
		}
		if !have{
			fmt.Println("没有这个字段，无法排序")
		}else {
			break
		}

	}

}



}