package todo

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

func ChoiceNew() {

	for {
		var choice string
		fmt.Println("请输入你的选择：1. 新增任务到管理系统  2. 查看管理系统中的任务")
		fmt.Scan(&choice)
		if choice == "1" {
			var (
				id, name, status, user string
			)
			for {
				fmt.Println("请输入要创建的任务ID，且任务ID不可以是 all :")
				if _, ok := msgTodo.todoItems[id]; ok {
					fmt.Println("任务ID已经存在，请重新输入!")
					break
				} else if id == "all" {
					fmt.Println("任务ID不可以是 \"all\" ,请重新输入!")
					break
				} else {
					fmt.Scan(&id)
					break
				}
			}

			fmt.Println("请输入任务名称:")
			fmt.Scan(&name)

			fmt.Println("请输入任务状态:")
			fmt.Scan(&status)

			fmt.Println("请输入任务执行者:")
			fmt.Scan(&user)
			msTodo := msgTodo.addTodo2msTodo(id, name, status, user)
			fmt.Println("请选择保存数据到文件的格式：1.gob 2.csv 3.json")
			var formatChoice string
			fmt.Scan(&formatChoice)
			if formatChoice == "1" {
				msTodo.gobSaveToFile("gobTask.txt")
			}else if formatChoice == "2" {
				msTodo.csvSaveToFile("csvTask.csv")
			}else if formatChoice == "3" {
				msTodo.jsonSaveToFile("jsonTask.json")
			}else {
				fmt.Println("输入错误，请重新选择！")
			}
		} else if choice == "2" {
			var id string
			fmt.Println("请输入任务ID：")
			fmt.Scan(&id)
			msgTodo.seeTodo2msTodo(id)
		}
	}
}



//新增todo到msTodo的todoItems中
func (msTodo *msTodo) addTodo2msTodo(id, name, status, user string) *msTodo {
	startTime := time.Now()
	endTime := startTime.Add(time.Hour * 24)
	//实例化一个todo
	todo := Newtodo(id, name, status, user, &startTime, &endTime)
	//将todo添加到任务管理系统中
	msTodo.todoItems[id] = append(msTodo.todoItems[id], todo)
	return msTodo
}

//查看任务
func (msTodo *msTodo) seeTodo2msTodo(id string) {
	//读取之前先清空msTodo
	for k, _ := range msTodo.todoItems {
		delete(msTodo.todoItems,k)
	}
	var choiceFormat string
	fmt.Println("请输入要查数据的文件格式：1. gob  2. csv  3. json")
	fmt.Scan(&choiceFormat)
	//调用函数从文件中读取数据
	if choiceFormat == "1" {
		msTodo.gobReadToFile("gobTask.txt")
	} else if choiceFormat == "2" {
		msTodo.csvReadToFile("csvTask.csv")
	} else if choiceFormat == "3" {
		msTodo.jsonReadToFile("jsonTask.json")
	} else {
		fmt.Println("输入错误，请重新选择！")
	}

	fmt.Println("++++++++++++++++++++++++++")

	//此时msgTodo.todoItems中已经有了数据
	if _,ok := msTodo.todoItems[id];ok {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"id","name","startTime","endTime","status","user"})
		for _,v := range msTodo.todoItems[id] {
				table.Append([]string{v.Id,v.Name,time2str(v.StartTime),time2str(v.EndTime),v.Status,v.User})
		}
		table.Render()
	}else {
		fmt.Println("任务ID不存在！")
	}
}
