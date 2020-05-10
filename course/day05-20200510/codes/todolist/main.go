package main

import (
	"fmt"
	"todolist/models"
)

func main() {
	var taskv2 models.Taskv2

	fmt.Printf("%#v\n", taskv2)

	taskv3 := models.Taskv3{
		Id:   1,
		Name: "完成Todolist",
	}

	taskv3.User = "kk"
	fmt.Printf("%#v\n", taskv3)

	anonyWrapperv4 := models.AnonyWrapperv4{}
	anonyWrapperv4.User = "kk"
	fmt.Printf("%#v\n", anonyWrapperv4)
}
