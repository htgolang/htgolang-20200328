package controllers

import "fmt"

type TaskController struct {
}

func (c *TaskController) Add() {
	panic("test")
	fmt.Println("添加任务成功")
}

func (c *TaskController) Query() {
	fmt.Println("查询成功")
}

func (c *TaskController) Delete() {
	fmt.Println("查询成功")
}

func (c *TaskController) Modify() {
	fmt.Println("修改成功")
}

func (c *TaskController) ModifyStatus() {
	fmt.Println("修改状态成功")
}
