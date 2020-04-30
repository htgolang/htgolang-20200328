package main

import (
	"fmt"
	"os"
	"time"
)
//任务存储在字典中，任务ID作为key,任务名称，任务计划开始时间，任务创建时间，任务状态为value
//声明并初始化一个字典
var taskMap = make(map[string][]string)
//新建任务
func createTask()  {
	var taskID,taskName,taskExecTime string
	fmt.Println("请输入任务ID：")
	fmt.Scan(&taskID)
	fmt.Println("请输入任务Name：")
	fmt.Scan(&taskName)
	fmt.Println("请输入任务计划执行时间：")
	fmt.Scan(&taskExecTime)
	CreteTime := time.Now().Format("2006-01-02 15:04")
	taskState := "未开始"
	taskMap[taskID] = []string{
		taskName,taskExecTime,CreteTime,taskState,
	}
}

//查看任务列表
func seeTask() {
	var taskName string
	fmt.Println("请输入任务名称或all来查看所有的任务信息:")
	fmt.Scan(&taskName)
	if taskName == "all" {
		if len(taskMap) == 0 {
			fmt.Println("暂无任务")
		}else {
			for k,v := range taskMap {
				fmt.Printf("任务ID：%s 任务名称：%s 任务计划执行时间：%s 任务创建时间：%s  任务状态：%s\n",k,v[0],v[1],v[2],v[3])
			}
		}
	}else {
		for k,v := range taskMap {
			if v[0] == taskName {
				fmt.Printf("任务ID：%s 任务名称：%s 任务计划执行时间：%s 任务创建时间：%s  任务状态：%s\n",k,v[0],v[1],v[2],v[3])
			} else {
				fmt.Println("任务不存在")
			}
		}
	}
}

//编辑任务
func editTask() {
	var ID string
	fmt.Println("开始编辑任务，请输入任务ID：")
	fmt.Scan(&ID)
	fmt.Printf("任务信息如下：\n 任务ID：%s  任务名称：%s  任务计划执行时间：%s  任务创建时间： %s  任务状态： %s\n", ID, taskMap[ID][0], taskMap[ID][1], taskMap[ID][2], taskMap[ID][3])
	var choice string
	fmt.Println("请输入Y/yes来选择是否继续编辑，输入其他则退出编辑：")
	fmt.Scan(&choice)
	if choice == "Y" || choice == "yes" {
		var taskName string
		var flag  = false
		for {
			fmt.Println("请输入任务名称：")
			fmt.Scan(&taskName)
			for _, v := range taskMap {
				if taskName == v[0] {
					fmt.Println("任务名称已经存在，请重新输入！")
					break
				} else {
					taskMap[ID][0] = taskName
					var editChoice string
					fmt.Println("请输入编辑的目的：\n 1：修改任务开始的计划时间\n 2: 修改任务状态：已完成、暂停、执行中")
					fmt.Scan(&editChoice)
					if editChoice == "1" {
						fmt.Println("请输入新的任务计划开始时间：例如：15:05")
						var execTaskTime string
						fmt.Scan(&execTaskTime)
						taskMap[ID][1] = execTaskTime
					}else if editChoice == "2" {
						fmt.Println("请输入任务状态：")
						var taskState string
						fmt.Scan(&taskState)
						if taskState == "已完成" {
							CreteTime := time.Now().Format("2006-01-02 15:04:05")
							taskMap[ID][3] = taskState+ ",完成时间："+CreteTime
						}else {
							taskMap[ID][3] = taskState
						}

					}
					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}
}

//删除任务
func deleteTask() {
	fmt.Println("请输入要删除的任务ID：")
	var taskID string
	fmt.Scan(&taskID)
	delete(taskMap, taskID)
}
func quitTask() {
	fmt.Println("退出程序！")
	os.Exit(-1)
}

func main() {
	var funcMap  = make(map[string]func())
	funcMap["1"] = createTask
	funcMap["2"] = seeTask
	funcMap["3"] = editTask
	funcMap["4"] = deleteTask
	funcMap["5"] = quitTask

	for {
		fmt.Println("请输入你的选择：\n 1：创建任务  2：查看任务  3：编辑任务 4: 删除任务  5：退出")
		var myChoice string
		fmt.Scan(&myChoice)
		if _,ok := funcMap[myChoice];ok {
			funcMap[myChoice]()
		}else {
			fmt.Println("输入错误，退出！")
			os.Exit(-1)
		}
	}
}