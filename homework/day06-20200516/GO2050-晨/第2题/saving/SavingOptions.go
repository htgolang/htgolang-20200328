package saving

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"todolist/task"
	"todolist/operations"
)
const TimeLayout = "2006-01-02 15:04:05"

func time2str(t *time.Time)  string{
	if t == nil {
		return ""
	}
	return t.Format(TimeLayout)
}
func Saving_option(tasks []*task.Task)  {
	text := operations.Input("请输入保存方式(gob/csv/json):")
	switch text {
	case "gob": GobSave(tasks)
	case "csv": CsvSave(tasks)
	case "json": JsonSave(tasks)
	default:
		fmt.Println("无效的保存选项!")
	}
}
func GobSave(tasks []*task.Task)  {
	file,err := os.Create("tasks.gob")
	if err != nil {
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(tasks)
}

func CsvSave(tasks []*task.Task)  {
	file,err := os.Create("tasks.csv")
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	csvWriter := csv.NewWriter(writer)
	csvWriter.Write([]string{"ID","任务名","开始时间","结束时间","任务状态","用户名字","用户地址","用户电话"})
	for _,task := range tasks{
		csvWriter.Write([]string{
			strconv.Itoa(task.GetIdTask()),
			task.GetNameTask(),
			time2str(task.GetstartTimeTask()),
			time2str(task.GetendTimeTask()),
			strconv.Itoa(task.GetStatusTask()),
			task.User.GetUserName(),
			task.User.GetUserAddr(),
			task.User.GetUserTel(),

		})
	}
}

func JsonSave(tasks []*task.Task)  {
	//fmt.Printf("%#v !!!!!!!!!!!!!!!!!!!\n",tasks[0].GetNameTask())
	ctx,_ := json.Marshal(tasks)
	//fmt.Printf("%#v  ~~~~~~~~~~~~~~\n",string(ctx))
	//if err != nil {
	//	return
	//}
	var buffer bytes.Buffer
	json.Indent(&buffer,ctx,"","\t")
	//buffer.WriteTo(os.Stdout)

	file,_ := os.Create("tasks.json")
	defer file.Close()
	buffer.WriteTo(file)

	//file,_ = os.Create("tasks2.json")
	//defer file.Close()
	//file.Write(ctx)
}