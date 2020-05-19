package reading

import (
	"bufio"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"todolist/task"

	"todolist/operations"
)
const TimeLayout = "2006-01-02 15:04:05"
var (
	gobpath = "tasks.gob"
	csvpath = "tasks.csv"
	jsonpath = "tasks.json"
)
func Reading_option()   []*task.Task{
	var tasks  []*task.Task
	var text = operations.Input("请输入读取的方式(gob/csv/json):")

	switch text {
	case "gob": tasks = GobRead(gobpath)
	case "csv": tasks = CsvRead(csvpath)
	case "json": tasks = JsonRead(jsonpath)
	default:
		fmt.Println("无效的读取选项！")
	}
	return tasks
}
func GobRead(path string) []*task.Task {
	var tasks []*task.Task
	file,_ :=os.Open(path)
	defer file.Close()

	decoder := gob.NewDecoder(file)
	decoder.Decode(&tasks)
	
	for _,task :=range tasks {
		fmt.Println(task)
	}
	return tasks
}

func JsonRead(path string) []*task.Task {
	var tasks []*task.Task
	jsonTxt,_ := ioutil.ReadFile(path)

	json.Unmarshal(jsonTxt,tasks)
	return tasks
}

func ParseTask(nodes []string)  *task.Task{
	id,_ := strconv.Atoi(nodes[0])
	name := nodes[1]
	st,_ := time.Parse(TimeLayout,nodes[2])
	startTime := &st
	et,_ := time.Parse(TimeLayout,nodes[3])
	endTime := &et
	status,_ := strconv.Atoi(nodes[4])
	username := nodes[5]
	addr := nodes[6]
	tel := nodes[7]
	user := &task.User{
		Name: username,
		Addr: addr,
		Tel: tel,
	}
	return &task.Task{
		Id: id,
		Name: name,
		StartTime: startTime,
		EndTime: endTime,
		Status: status,
		User: user,
	}
}

func CsvRead(path string)  []*task.Task{
	var tasks []*task.Task
	file,_ := os.Open(path)
	defer file.Close()

	reader := bufio.NewReader(file)
	csvReader := csv.NewReader(reader)
	csvReader.Read()

	for {
		line,err1 := csvReader.Read()
		if err1 == io.EOF{
			break
		}
		task := ParseTask(line)

		tasks = append(tasks,task)
	}
	return tasks
}