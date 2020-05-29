package controllers

import (
	"encoding/json"
	"fmt"
	"myTodolist/utils/ioutils"
	"os"
	"time"
)

//定义一个任务的结构体
type TaskController struct {
	Id string
	Name string
	StartTime *time.Time
	EndTime *time.Time
	Status string
	User string
}

var TaskList = []*TaskController{}

//定义一个添加任务的方法
/*
为了解决已经写入到json文件中的内容在下次运行程序时依然保存，需要验证文件是否存在
	1. 已经存在的话：需要先读取json内容到TaskList中，然后再添加Task，最后保存(保存之前清空文件)到json文件中
	2. 不存在的话，则直接打开文件
 */
func (c *TaskController) AddTask() {
	if _,err := os.Stat("taskJson.json");err != nil {
		//如果错误是文件不存在
		if os.IsNotExist(err) {
			//创建并打开一个文件
			file, err := os.OpenFile("taskJson.json",os.O_CREATE|os.O_RDWR,0644)
			defer file.Close()
			//创建不成功的话打印出错误并返回
			if err != nil {
				ioutils.Error(fmt.Sprintf(err.Error()))
				return
			}
			//用户输入，数据验证，添加数据
			id := ioutils.Input("请输入任务ID：")
			name := ioutils.Input("请输入任务名称：")
			status := ioutils.Input("请输入任务状态")
			user := ioutils.Input("请输入任务执行者")
			task := NewTaskController(id,name,status,user)
			TaskList = append(TaskList,task)

			encodeJson := json.NewEncoder(file)
			encodeJson.Encode(TaskList)
		}else { //否则是其他错误
			if err != nil {
				ioutils.Error(fmt.Sprintf(err.Error()))
				return
			}
		}
	}else {//否则文件已经存在
		//文件已经存在的话则Decode文件内容到TaskList中
		c.SeeTask()
		//然后清空文件内的内容
		file, err := os.Create("taskJson.json")
		defer file.Close()
		if err != nil {
			ioutils.Error(fmt.Sprintf(err.Error()))
			return
		}
		//用户输入，数据验证，添加数据
		id := ioutils.Input("请输入任务ID：")
		name := ioutils.Input("请输入任务名称：")
		status := ioutils.Input("请输入任务状态")
		user := ioutils.Input("请输入任务执行者")
		task := NewTaskController(id,name,status,user)
		TaskList = append(TaskList,task)

		encodeJson := json.NewEncoder(file)
		encodeJson.Encode(TaskList)

	}

}

//封装一个函数，返回TaskController结构体的类型指针
func NewTaskController(id, name, status, user string) *TaskController {
	nowTime := time.Now()
	endTime := nowTime.Add(time.Hour*24)
	return &TaskController{
		Id: id,
		Name: name,
		StartTime: &nowTime,
		EndTime: &endTime,
		Status: status,
		User: user,
	}
}