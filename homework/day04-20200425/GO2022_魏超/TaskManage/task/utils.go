package task

import "fmt"

// SaveDB 保存数据
func SaveDB(tasksInfo map[int]*TaskType, dbFile string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("新建任务失败,%s", err)
		}
	}()
	defer func() {
		err := WriteDB(tasksInfo, dbFile)
		if err != nil {
			panic(err)
		}
	}()
}

// GenerateTaskID 生成任务信息的ID
func GenerateTaskID(tasksInfo map[int]*TaskType) int {
	var rt int
	for _, task := range tasksInfo {
		id := task.ID
		if rt < id {
			rt = id
		}
	}
	return rt + 1
}
