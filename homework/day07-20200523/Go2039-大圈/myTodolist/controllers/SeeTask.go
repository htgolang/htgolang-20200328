package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"myTodolist/utils/ioutils"
	"os"
	"time"
)

func time2String(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

//定义一个查看任务的方法
func (c *TaskController) SeeTask() {
	//打开json文件
	file, err := os.Open("taskJson.json")
	if err != nil {
		ioutils.Error(fmt.Sprintf(err.Error()))
		return
	}
	jsonDecode := json.NewDecoder(file)
	jsonDecode.Decode(&TaskList)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id","Name","StartTime","EndTime","Status","User"})
	for i := 0; i<len(TaskList); i++ {
		table.Append([]string{
			TaskList[i].Id,
			TaskList[i].Name,
			time2String(TaskList[i].StartTime),
			time2String(TaskList[i].EndTime),
			TaskList[i].Status,
			TaskList[i].User},
			)
	}
	table.Render()
}
