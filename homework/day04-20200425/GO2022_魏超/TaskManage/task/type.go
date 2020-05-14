package task

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shadow_wei/TaskManage/user"
	"github.com/shadow_wei/TaskManage/utils"
)

// TaskType 用户类型
type TaskType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartTime int64  `json:"start"`
	EndTime   int64  `json:"endtime"`
	Status    int    `json:"status"`
	UserID    int    `json:"uid"`
}

var TaskStat = map[int]string{
	0: "新创建",
	1: "已完成",
	2: "开始执行",
	3: "暂停",
}

// New 创建taskType
func (task TaskType) New(id, uid int, name string, starttime int64) *TaskType {
	return &TaskType{
		ID:        id,
		Name:      name,
		StartTime: starttime,
		Status:    0,
		UserID:    uid,
	}
}

// IsEqual 根据属性名和值判断是否相等
func (task TaskType) IsEqual(field, value string) bool {
	switch strings.ToLower(field) {
	case "id":
		id, err := strconv.Atoi(value)
		if err == nil {
			if task.ID == id {
				return true
			}
		}
	case "name":
		return strings.Contains(task.Name, value)
	case "StartTime":
		starttime, err := time.Parse("2006-01-02 15:04:05", value)
		if err != nil {
			return false
		}
		return time.Unix(task.StartTime, 0).Truncate(1 * time.Second).Equal(starttime)
	case "EndTime":
		endtime, err := time.Parse("2006-01-02 15:04:05", value)
		if err != nil {
			return false
		}
		return time.Unix(task.EndTime, 0).Truncate(1 * time.Second).Equal(endtime)
	case "Status":
		status, err := strconv.Atoi(value)
		if err == nil {
			if task.Status == status {
				return true
			}
		}
	case "UserID":
		uid, err := strconv.Atoi(value)
		if err == nil {
			if task.UserID == uid {
				return true
			}
		}
	default:
		fmt.Printf("您输入的%s字段是不正确.\n", field)
	}
	return false
}

// FormatSlice 转成切片格式
func (task TaskType) FormatSlice(format []string) []string {
	taskSlice := make([]string, 0)
	for _, filed := range format {
		switch strings.ToLower(filed) {
		case "id":
			taskSlice = append(taskSlice, strconv.Itoa(task.ID))
		case "name":
			taskSlice = append(taskSlice, task.Name)
		case "starttime":
			taskSlice = append(taskSlice, time.Unix(task.StartTime, 0).Format("2006-01-02 15:04:05"))
		case "endtime":
			taskSlice = append(taskSlice, time.Unix(task.EndTime, 0).Format("2006-01-02 15:04:05"))
		case "status":
			statName, _ := TaskStat[task.Status]
			taskSlice = append(taskSlice, statName)
		case "userid":
			taskSlice = append(taskSlice, user.GetUserName(task.UserID))
		default:
			taskSlice = append(taskSlice, "")
		}
	}
	return taskSlice
}

// ModifyTask 修改任务数据
func (task *TaskType) ModifyTask(field string) error {
	var err error
	switch strings.ToLower(field) {
	case "name":
		value := utils.Input("请输入任务名称：")
		if IsExitstTask(value) {
			err = errors.New("任务存在.")
		} else {
			task.Name = value
		}
	case "starttime":
		if starttime, err := time.Parse("2006-01-02 15:04:05", utils.Input("请输入开始时间(yyyy-mm-dd HH:MM:SS)：")); err != nil {
			fmt.Printf("输入的时间格式非法，%s，创建任务失败.\n", err)
		} else {
			task.StartTime = starttime.Unix()
		}
	case "status":
		fmt.Println("可以选择的任务状态：")
		for _, statusname := range TaskStat {
			fmt.Println(statusname)
		}
		value := utils.Input("请输入任务状态：")
		for statusID, statusName := range TaskStat {
			if statusName == value {
				task.Status = statusID
				if task.Status == 1 {
					task.EndTime = time.Now().Unix()
				}
				break
			}
		}
	default:
		err = errors.New(fmt.Sprintf("您输入的%s字段是不正确.", field))
	}
	return err
}
