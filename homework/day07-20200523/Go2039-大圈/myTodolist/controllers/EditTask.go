package controllers

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"myTodolist/utils/ioutils"
	"os"
)

//定义一个编辑任务的方法
func (c *TaskController) EditTask() {
	//先读取内容到TaskList中
	c.SeeTask()
	//提示用户输入要编辑的任务ID
	id := ioutils.Input("请输入要编辑的任务ID：")
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
				name := ioutils.Input("请输入Name")
				TaskList[i].Name = name
				user := ioutils.Input("请输入User")
				TaskList[i].User = user
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
