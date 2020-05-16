package main

import (
	"bytes"
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

	// json编码
	ctx, _ := json.Marshal(tasks)

	// 将json格式化输出到buffer对象
	var buffer bytes.Buffer
	json.Indent(&buffer, ctx, "", "\t")

	//将buffer对象输出到标准输出
	// buffer.WriteTo(os.Stdout)

	// 创建文件，并将格式化的字符串输出到json文件
	file, _ := os.Create("task.json")
	defer file.Close()
	buffer.WriteTo(file)

	// 创建文件，并将json格式字符串输出到json文件
	file2, _ := os.Create("task2.json")
	defer file2.Close()
	file2.Write(ctx)

}
