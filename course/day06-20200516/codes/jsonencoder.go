package main

import (
	"encoding/json"
	"os"
	"time"
)

// 定义task结构体
type Task struct {
	Id        int
	Name      string
	Status    int
	StartTime *time.Time
	EndTime   *time.Time
	User      string
}

func main() {
	now := time.Now()
	end := now.Add(time.Hour * 24)

	tasks := []*Task{
		{
			Id:        1,
			Name:      "整理课程笔记",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "kk",
		},
		{
			Id:        2,
			Name:      "整理课程笔记2",
			Status:    0,
			StartTime: &now,
			EndTime:   &end,
			User:      "kk",
		},
	}

	file, _ := os.Create("task.json")

	defer file.Close()

	// 创建json编码器
	encoder := json.NewEncoder(file)

	// 编码内容到文件中
	encoder.Encode(tasks)

}
