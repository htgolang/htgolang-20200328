package task

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"todolistv3/utils"

	"github.com/olekukonko/tablewriter"
)

const (
	statusNew      = "未执行"
	statusComplete = "完成"
	statusBegin    = "开始执行"
	statusPause    = "暂停"
)

type Task struct {
	ID        int
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Status    string
	User      string
}

var (
	statusChoice = []string{statusNew, statusBegin, statusComplete, statusPause}
	header       = []string{"ID", "Name", "StartTime", "EndTime", "Status", "User"}
)

// Task结构体方法
func (task *Task) SetName(name string) {
	task.Name = name
}

func (task *Task) SetStatus(status string) {
	task.Status = status
}

func (task *Task) SetUser(user string) {
	task.User = user
}

func (task *Task) SetStartTime(startTime time.Time) {
	task.StartTime = startTime
}

func (task *Task) SetEndtime(endtime time.Time) {
	task.EndTime = endtime
}

// Task结构体初始化函数
func newTask(tasks []Task) *Task {
	id := genID(tasks)
	now := time.Now()
	end := now.Add(24 * time.Hour)
	return &Task{
		ID:        id,
		Name:      "",
		StartTime: now,
		EndTime:   end,
		Status:    statusNew,
		User:      "",
	}
}

// 生成tasks最大的id
func genID(tasks []Task) int {
	var rt int
	if len(tasks) == 0 {
		return 1
	}

	for _, task := range tasks {
		if rt < task.ID {
			rt = task.ID
		}
	}
	return rt + 1
}

// 验证状态值在可选列表内
func verifyStatus(inputStatus string) bool {
	for _, status := range statusChoice {
		if inputStatus == status {
			return true
		}
	}
	return false
}

// 验证任务名，确保唯一性
func verifyName(tasks []Task, inputName string) bool {
	for _, task := range tasks {
		if inputName == task.Name {
			return false
		}
	}
	return true
}

//渲染输出任务信息
func RenderTask(tasks ...Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(header)
	for i := 0; i < len(header); i++ {
		table.SetColMinWidth(i, 20)
	}
	for _, task := range tasks {
		table.Append([]string{
			strconv.Itoa(task.ID),
			task.Name,
			task.StartTime.Format("2006/01/02 15:04:05"),
			task.EndTime.Format("2006/01/02 15:04:05"),
			task.Status,
			task.User,
		})
	}
	table.Render()
	// table.AppendBulk(content)
}

// 添加一个任务
func AddTask(tasks []Task, path string) Task {
	task := newTask(tasks)
	fmt.Println("请输入任务信息!")

	for {
		tempName := utils.Input("任务名:")
		if verifyName(tasks, tempName) {
			task.SetName(tempName)
			break
		} else {
			fmt.Println("任务名称已存在!")
		}
	}
	task.SetUser(utils.Input("负责人:"))

	utils.ChangePassword(path)

	RenderTask(*task)
	return *task
}

// 排序任务信息
func sortTask(tasks []Task, key string) []Task {
	if key == "name" {
		sort.Slice(tasks, func(i, j int) bool { return tasks[i].Name < tasks[j].Name })
	}
	if key == "startTime" {
		sort.Slice(tasks, func(i, j int) bool { return tasks[i].StartTime.Before(tasks[j].StartTime) })
	}
	return tasks
}

// 查询任务
func QueryTaskWithSort(tasks []Task) {
	queryMap := map[string]string{
		"1": "name",
		"2": "startTime",
	}
	if len(tasks) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
		return
	}
	q := utils.Input("请输入查询的任务名称:")
	taskContent := make([]Task, 0)

	for _, task := range tasks {
		if q == "all" || strings.Contains(task.Name, q) {
			taskContent = append(taskContent, task)
		}
	}
	if len(taskContent) == 0 {
		fmt.Println("未找到关联任务!")
	} else {
		key := utils.Input("请输入排序方式[1.任务名称 2.任务开始时间]:")
		newTasks := sortTask(taskContent, queryMap[key])
		RenderTask(newTasks...)
	}
}

// 修改任务
func ModifyTask(tasks []Task, path string) {
	if len(tasks) == 0 {
		fmt.Println("目前没有任何任务记录，请先添加任务，谢谢!")
		return
	}
	q := utils.Input("请输入需要修改的任务ID:")
	for index, task := range tasks {
		if id, _ := strconv.Atoi(q); id == task.ID {
			RenderTask(task)
			switch utils.Input("是否修改,请确认(y/yes):") {
			case "y", "yes":
				for {
					tempName := utils.Input("任务名称:")
					if verifyName(tasks, tempName) {
						task.SetName(tempName)
						break
					} else {
						fmt.Println("任务名称已存在!")
					}
				}
				for {
					tempStatus := utils.Input("状态:")
					if verifyStatus(tempStatus) {
						if tempStatus == statusComplete {
							task.SetEndtime(time.Now())
						}
						task.SetStatus(tempStatus)
						break
					} else {
						fmt.Println("输入的状态值不对!可选范围:", strings.Join(statusChoice, ", "))
					}
				}
				tasks[index] = task
				utils.ChangePassword(path)
				fmt.Println("任务修改完成!")
				RenderTask(task)
			default:
				fmt.Println("取消修改!")
				break
			}
		}
	}
}

// 删除任务
func DeleteTask(tasks []Task, path string) []Task {
	newTasks := make([]Task, 0)
	queryId := utils.Input("请输入需要删除的任务ID:")
	for index, task := range tasks {
		if qid, _ := strconv.Atoi(queryId); qid == task.ID {
			RenderTask(task)
			switch utils.Input("是否进行删除(y/yes):") {
			case "y", "yes":
				copy(tasks[index:], tasks[index+1:])
				newTasks = tasks[:len(tasks)-1]
				utils.ChangePassword(path)
				fmt.Printf("任务ID:%s 已删除\n", queryId)
			default:
				fmt.Println("取消删除!")
			}
		}
	}
	return newTasks
}

// 持久化任务信息
func RecordTask(path string, tasks ...Task) {
	content := make([]string, 0)
	for _, task := range tasks {
		taskContent := []string{
			strconv.Itoa(task.ID),
			task.Name,
			task.StartTime.Format("2006/01/02 15:04:05"),
			task.EndTime.Format("2006/01/02 15:04:05"),
			task.Status,
			task.User,
		}
		content = append(content, strings.Join(taskContent, ","))
	}
	utils.WriteFile(path, strings.Join(content, "\n"))
}

// 从文件读取task信息
func ReadTaskFromFile(path string) []Task {
	var taskRecords []Task
	var task Task

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		taskSlice := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(taskSlice[0])
		startTime, _ := time.Parse("2006/01/02 15:04:05", taskSlice[2])
		endTime, _ := time.Parse("2006/01/02 15:04:05", taskSlice[3])
		task = Task{id, taskSlice[1], startTime, endTime, taskSlice[4], taskSlice[5]}
		taskRecords = append(taskRecords, task)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return taskRecords
}
