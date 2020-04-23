package main

import (
	"fmt"
	"time"
)

/*
4. todo管理
	1. 新建任务（任务ID，任务名称，任务创建时间，任务计划开始时间，任务目前的状态）
	2. 查看任务（查看所有的任务／查看单个任务）

    a. 编辑
        请输入编辑的ID:
        通过ID查找 => Task => 显示
        用户确认是否进行编辑(y/yes): 编辑
        用户输入: 任务名称, 开始时间，状态
        // 状态如果是已完成, 初始化完成时间
        // time.Now().Format("2006-01-02 15:04:05")

    b. 删除
        请输入编辑的ID:
        通过ID查找 => Task => 显示
        用户确认是否进行删除(y/yes): 删除

    c. 数据验证（用户输入数据检查）
        任务名称: 不能重复 (新增，编辑)
        编辑：
            a. 任务名称不变的时候则正常编辑
            b. 任务名称改成其他已存在的任务名称时候，提示任务名称已存在，提示重新输入任务名称或者退出编辑，再次选择执行操作
        任务状态:
            未开始, 执行中, 已完成, 暂停
 */
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
		for k,v := range taskMap {
			fmt.Printf("任务ID：%s 任务名称：%s 任务计划执行时间：%s 任务创建时间：%s  任务状态：%s\n",k,v[0],v[1],v[2],v[3])
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
					v[0] = taskName
					var editChoice string
					fmt.Println("请输入编辑的目的：\n 1：修改任务开始的计划时间\n 2: 修改任务状态：已完成、暂停、执行中")
					fmt.Scan(&editChoice)
					if editChoice == "1" {
						fmt.Println("请输入新的任务计划开始时间：例如：15:05")
						var execTaskTime string
						fmt.Scan(&execTaskTime)
						v[1] = execTaskTime
					}else if editChoice == "2" {
						fmt.Println("请输入任务状态：")
						var taskState string
						fmt.Scan(&taskState)
						if taskState == "已完成" {
							CreteTime := time.Now().Format("2006-01-02 15:04")
							v[3] = taskState+ ",完成时间："+CreteTime
						}else {
							v[3] = taskState
						}

					}
					flag = true
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

func main() {
	for {
		fmt.Println("请输入你的选择：\n 1：创建任务  2：查看任务  3：编辑任务 4: 删除任务  5：退出")
		var myChoice string
		fmt.Scan(&myChoice)
		if myChoice == "1" {
			createTask()
		} else if myChoice == "2" {
			if len(taskMap) == 0  {
				fmt.Println("暂无任务")
			}else {
				seeTask()
			}
		} else if myChoice == "3" {
			editTask()
		} else if myChoice == "5" {
			break
		} else if myChoice == "4" {
			deleteTask()
		} else {
			fmt.Println("输入错误，请重新输入！")
		}
	}
}