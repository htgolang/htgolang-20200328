package main

import (
	"encoding/json"
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

func main() {
	var tasks []*Task

	file, _ := os.Open("task.json")
	decoder := json.NewDecoder(file)

	decoder.Decode(&tasks)

	for _, task := range tasks {
		fmt.Printf("%#v\n", task)
	}

}
