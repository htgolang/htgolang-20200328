package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shadow_wei/TaskManage/user"
	"github.com/shadow_wei/TaskManage/utils"
)

// tasksInfo 声明定义任务
var tasksInfo map[int]*TaskType

var dbFile = "taskinfo.db"
var taskFormat = []string{"id", "name", "status", "starttime", "endtime", "userid"}

func init() {
	execuFile, err := os.Executable()
	if err != nil {
		panic(fmt.Sprintf("获取当前执行目录失败,%s\n", err))
	} else {
		dbFile = path.Join(filepath.Dir(execuFile), dbFile)
	}
	if dbInfo := []byte(ReadDB(dbFile)); len(dbInfo) > 0 {
		err = json.Unmarshal(dbInfo, &tasksInfo)
		if err != nil {
			panic(fmt.Sprintf("无法从%s反序列化数据,请检查数据的有效性.err:%s", dbFile, err))
		}
	} else {
		tasksInfo = make(map[int]*TaskType)
	}
}

// NewTask 创建新任务
func NewTask() {
	var err error
	var task TaskType
	id := GenerateTaskID(tasksInfo)
	name := utils.Input("请输入任务名：")
	if IsExitstTask(name) {
		fmt.Println("存在相同任务")
		return
	}
	starttime, err := time.Parse("2006-01-02 15:04:05", utils.Input("请输入开始时间(yyyy-mm-dd HH:MM:SS)："))
	if err != nil {
		fmt.Printf("输入的时间格式非法，%s，创建任务失败.\n", err)
		return
	}
	implementer := utils.Input("请输入责任人：")
	uid, err := user.GetUserID(implementer)
	if err != nil {
		fmt.Println("责任人不存在.")
		return
	}

	tasksInfo[id] = TaskType.New(task, id, uid, name, starttime.Unix())
	SaveDB(tasksInfo, dbFile)
}

// QueryTask 查询任务
func QueryTask() {
	taskSlice := make([][]string, 0)
	fmt.Println("字段信息：\nID\nName\nStartTime\nEndTime\nStatus\nUserID")
	field := utils.Input("请输入要查询的字段名称：")
	value := utils.Input("请输入查询的信息：")

	for _, task := range tasksInfo {
		if value == "all" || task.IsEqual(field, value) {
			taskSlice = append(taskSlice, task.FormatSlice(taskFormat))
		}
	}
	switch strings.ToLower(field) {
	case "id":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][0] < taskSlice[j][0] })
	case "name":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][1] < taskSlice[j][1] })
	case "Status":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][2] < taskSlice[j][2] })
	case "StartTime":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][3] < taskSlice[j][3] })
	case "EndTime":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][4] < taskSlice[j][4] })
	case "UserID":
		sort.Slice(taskSlice, func(i, j int) bool { return taskSlice[i][5] < taskSlice[j][5] })
	}
	utils.TableFormat(taskFormat, taskSlice)
}

// DeleteTask 删除任务
func DeleteTask() {
	taskSlice := make([][]string, 0)
	taskID, err := strconv.Atoi(utils.Input("请输入需要删除的任务ID："))
	if err != nil {
		fmt.Println("输入的任务ID格式不正确.")
		return
	}
	task, ok := tasksInfo[taskID]
	if ok == true {
		taskSlice = append(taskSlice, task.FormatSlice(taskFormat))
	}
	utils.TableFormat(taskFormat, taskSlice)
	if isDelete := utils.Input("是否确认删除当前任务信息(y/yes):"); isDelete == "yes" || isDelete == "y" {
		delete(tasksInfo, taskID)
	}
	SaveDB(tasksInfo, dbFile)
}

// ModifyTask 修改任务
func ModifyTask() {
	taskSlice := make([][]string, 0)
	taskID, err := strconv.Atoi(utils.Input("请输入需要编辑的任务ID："))
	if err != nil {
		fmt.Println("输入的任务ID格式不正确.")
		return
	}
	task, ok := tasksInfo[taskID]
	if ok == false {
		fmt.Println("输入的任务ID不存在.")
		return
	}
	taskSlice = append(taskSlice, task.FormatSlice(taskFormat))
	utils.TableFormat(taskFormat, taskSlice)

	if isDelete := utils.Input("是否确认修改当前任务信息(yes/y):"); isDelete == "yes" || isDelete == "y" {
		for {
			fmt.Println("字段信息：\nName\nStartTime\nStatus")
			field := utils.Input("请输入您要修改的字段名称：")
			err := task.ModifyTask(field)
			if err != nil {
				fmt.Println("修改数据失败,", err)
			} else {
				SaveDB(tasksInfo, dbFile)
				fmt.Println("修改成功.")
			}
			if isExit := utils.Input("是否确定继续修改此任务信息(yes/y):"); !(isExit == "yes" || isExit == "y") {
				break
			}
		}
	}
}

// IsExitstTask 检测用户的账号是否存在
func IsExitstTask(tname string) bool {
	for _, task := range tasksInfo {
		if task.Name == tname {
			return true
		}
	}
	return false
}
