package controllers

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"myTodolist/utils/ioutils"
	"os"
)

//定义一个修改任务状态的方法
func (c *TaskController) EditTaskStatus() {
	//读取文件的内容到TaskList中
	c.SeeTask()
	id := ioutils.Input("请输入要修改任务状态的任务ID：")
	//打印出用户选择的任务，让用户确认
	for i:=0;i<len(TaskList);i++ {
		if TaskList[i].Id == id {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Id","Name","StartTime","EndTime","Status","User"})
			table.Append([]string{
				TaskList[i].Id,
				TaskList[i].Name,
				time2String(TaskList[i].StartTime),
				time2String(TaskList[i].EndTime),
				TaskList[i].Status,
				TaskList[i].User},
			)
			table.Render()
			confirm := ioutils.Input("请确认任务ID y or yes：")
			if confirm == "yes" || confirm == "y" {
				status := ioutils.Input("请输入任务状态")
				TaskList[i].Status = status
				//当文件存在时则清空文件内容
				file, _ := os.Create("taskJson.json")
				defer file.Close()
				jsonEncode := json.NewEncoder(file)
				jsonEncode.Encode(TaskList)
				break
			}else {
				ioutils.Error("输入错误")
				break
			}
		}
		if i == len(TaskList)-1 {
			ioutils.Error("任务ID不存在！")
		}
	}

}
