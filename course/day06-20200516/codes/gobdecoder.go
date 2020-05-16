package main

import (
	"encoding/gob"
	"fmt"
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
	// 注册对象到gob管理器中
	gob.Register(&Task{})
}

func main() {

	var tasks []*Task

	file, _ := os.Open("user.gob")

	defer file.Close()

	// 定义解码器
	decoder := gob.NewDecoder(file)

	// 解码
	decoder.Decode(&tasks)

	for _, task := range tasks {
		fmt.Printf("%T\n", task)
	}
}
