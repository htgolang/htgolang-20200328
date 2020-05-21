package task

import (
	"bufio"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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
//数据存储方式
var FileMode = "default"

type Task struct {
	Id int
	Name string
	StartTime time.Time
	EndTime time.Time
	Status string
	User string
	Del bool
}

var (
	LoadTask func()
	WriteTask func()
)

//从用户输入数据存储方式
func UserDataType()  {
	userinput := auth.Input("请输入数据存储方式\n1:default\n2:json\n3:csv\n4:gob\n")
	switch userinput {
	case "1", "default":
		LoadTask = LoadTaskSelf
		WriteTask =WriteTaskSelf
	case "2", "json":
		LoadTask = LoadTaskJson
		WriteTask =WriteTaskJson
	case "3", "csv":
		LoadTask = LoadTaskCsv
		WriteTask =WriteTaskCsv
	case "4", "gob":
		LoadTask = LoadTaskGob
		WriteTask =WriteTaskGob
	default:
		LoadTask = LoadTaskSelf
		WriteTask =WriteTaskSelf
	}
}

//string ->task
func parseTask(node string)*Task  {
	tmpSlice := strings.Split(node,",")
	id,_ := strconv.Atoi(tmpSlice[0])
	name :=tmpSlice[1]
	startTime,_ := time.Parse(Layout,tmpSlice[2])
	endTime,_ := time.Parse(Layout,tmpSlice[3])
	status := tmpSlice[4]
	user := tmpSlice[5]
	del,_ := strconv.ParseBool(tmpSlice[6])
	task :=Task{
		Id: id,
		Name:name,
		StartTime:startTime,
		EndTime:endTime,
		Status:status,
		User:user,
		Del:del,
	}
	return &task
}
//task ->string
func deParseTask(task *Task) string  {
	id := strconv.Itoa(task.Id)
	name := task.Name
	startTime := task.StartTime.Format(Layout)
	endTime := task.EndTime.Format(Layout)
	status := task.Status
	user := task.User
	del := strconv.FormatBool(task.Del)
	strLin := strings.Join([]string{id,name,startTime,endTime,status,user,del},",")
	return strLin
}

func LoadTaskSelf()  {
	file,_ := os.OpenFile("task.txt",os.O_RDONLY,os.ModePerm)
	defer file.Close()
	scaner := bufio.NewScanner(file)
	for scaner.Scan(){
		tmp :=scaner.Text()
		task :=parseTask(tmp)
		Todolist = append(Todolist,*task)
	}

}

func WriteTaskSelf()  {

	file,_ := os.OpenFile("task.txt",os.O_RDONLY|os.O_TRUNC|os.O_CREATE,os.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _,task :=range Todolist{

		strLin := deParseTask(&task)
		writer.WriteString(strLin+"\n")
		writer.Flush()
	}
	backFile,_ := os.OpenFile(filepath.Join("back",CheckBackFileName("task.txt")),os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	defer backFile.Close()
	file.Seek(0,0)
	copyFile(file,backFile)
}
//加载文件数据到todolist
func LoadTaskJson()  {
	file,_ := os.OpenFile("task.json",os.O_CREATE|os.O_RDONLY,os.ModePerm)
	defer file.Close()
	buffReader := bufio.NewReader(file)
	jsondecode := json.NewDecoder(buffReader)
	jsondecode.Decode(&Todolist)

}
//加载数据todolist到文件
func WriteTaskJson()  {
	file,_ := os.OpenFile("task.json",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,os.ModePerm)
	defer file.Close()
	buffwriter := bufio.NewWriter(file)
	jsonEncode :=json.NewEncoder(buffwriter)
	jsonEncode.SetIndent("","\t")
	jsonEncode.Encode(Todolist)
	defer buffwriter.Flush()

	backFile,_ := os.OpenFile(filepath.Join("back",CheckBackFileName("task.json")),os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	defer backFile.Close()
	file.Seek(0,0)
	copyFile(file,backFile)

}
func LoadTaskCsv()  {
	file,_ := os.OpenFile("task.csv",os.O_CREATE|os.O_RDONLY,os.ModePerm)
	defer file.Close()
	buffReader := bufio.NewReader(file)
	csvReader := csv.NewReader(buffReader)
	csvReader.Read()
	for {
		line,err := csvReader.Read()
		if err !=nil{
			if err == io.EOF{
				break
			}
			fmt.Println(line,err)
			break
		}
		task :=parseTask(strings.Join(line,","))
		Todolist = append(Todolist,*task)

	}

}
func WriteTaskCsv()  {
	file,_ := os.OpenFile("task.csv",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,os.ModePerm)
	defer file.Close()
	buffWriter := bufio.NewWriter(file)
	csvWriter := csv.NewWriter(buffWriter)
	csvWriter.Write([]string{"Id","Name","StartTime","EndTime","Status","User","Del"})
	for _,j := range Todolist{
		csvWriter.Write(strings.Split(deParseTask(&j),","))

	}
	defer buffWriter.Flush()

	backFile,_ := os.OpenFile(filepath.Join("back",CheckBackFileName("task.csv")),os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	defer backFile.Close()
	file.Seek(0,0)
	copyFile(file,backFile)
}
func LoadTaskGob()  {
	file,_ := os.OpenFile("task.gob",os.O_CREATE|os.O_RDONLY,os.ModePerm)
	defer file.Close()
	buffReader := bufio.NewReader(file)
	gobNewReader := gob.NewDecoder(buffReader)
	gobNewReader.Decode(&Todolist)
}
func WriteTaskGob()  {
	file,_ := os.OpenFile("task.gob",os.O_CREATE|os.O_WRONLY|os.O_TRUNC,os.ModePerm)
	defer file.Close()
	buffWriter := bufio.NewWriter(file)
	gobNewWriter := gob.NewEncoder(buffWriter)
	gobNewWriter.Encode(Todolist)
	defer buffWriter.Flush()

	backFile,_ := os.OpenFile(filepath.Join("back",CheckBackFileName("task.gob")),os.O_CREATE|os.O_TRUNC|os.O_WRONLY,os.ModePerm)
	defer backFile.Close()
	file.Seek(0,0)
	copyFile(file,backFile)
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
			if j.Del == true{
				fmt.Println("没有这个id")
				break
			}
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
			if j.Del == true{
				fmt.Println("没有这个id")
				break
			}
			show(j)
			userInPut := auth.Input("以上是任务信息，确认要删除吗？(yes/y)")
			switch userInPut {
			case "yes","y":
				Todolist[i].Del = true
			default:

			}
			break
		}
		if i == len(Todolist)-1{
			fmt.Println("没有这个id")
		}
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
				if j.Del == true{
					fmt.Println("没有这个id")
					break
				}
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
	WriteTask()
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
		if j.Del{
			continue
		}
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

//备份数据,根据时间排序，删除最老的备份
func CheckBackFileName(prefix string)(fileName string){

	type fileTime struct{
		name string
		newTime time.Time
	}

	fileSlice := make([]fileTime,0)
	_,err := os.Stat("back")
	if err !=nil{
		if os.IsNotExist(err){
			os.Mkdir("back",os.ModePerm)
		}
	}

	filepath.Walk("back", func(path string, info os.FileInfo, err error) error {
		if match,err:=regexp.Match(prefix,[]byte(info.Name())); err == nil{
			if match{
				fileSlice = append(fileSlice, fileTime{info.Name(),info.ModTime()})
				info.ModTime()
			}
		}
		return nil
	})

	if len(fileSlice)<3{
		file,_ := ioutil.TempFile("back",prefix)
		defer file.Close()
		return file.Name()
	}else {
		sort.SliceStable(fileSlice, func(i, j int) bool {
			return fileSlice[i].newTime.Before(fileSlice[j].newTime)
		})
		return fileSlice[0].name
	}
}
//copyfile
func copyFile(src io.Reader, dest io.Writer)  {
	io.Copy(dest,src)
}