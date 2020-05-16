package main

import (
	"encoding/gob"
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

func init() {
	// 注册持久化的对象到gob管理器中
	gob.Register(&Task{})
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

	// 创建文件
	file, _ := os.Create("user.gob")

	defer file.Close()

	// 创建encode对象
	encoder := gob.NewEncoder(file)

	// encode内存中对象到文件
	encoder.Encode(tasks)
}
