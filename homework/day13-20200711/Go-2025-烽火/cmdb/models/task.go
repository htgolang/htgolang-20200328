package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// const (
// 	TimeLayout = "2006-01-02T15:04:05"
// 	DateLayout = "2006-01-02 15:04:05"
// )

type Task struct {
	ID           int        `form:"id" orm:"column(id)"`
	Name         string     `form:"name" orm:"size(64)"`
	StartTime    *time.Time `form:"start_time" orm:"column(start_time)"`
	CompleteTime *time.Time `form:"complete_time" orm:"column(complete_time);null"`
	DeadlineTime *time.Time `form:"deadline_time" orm:"column(deadline_time);"`
	Status       int        `form:"status" orm:""`
	Content      string     `form:"content" orm:"null"`
	User         int        `form:"user" orm:""`
}

func init() {
	orm.RegisterModel(new(Task))
}

func (task *Task) StatusText() string {
	switch task.Status {
	case 0:
		return "新建"
	case 1:
		return "开始"
	case 2:
		return "暂停"
	case 3:
		return "完成"
	}
	return "未知"
}

func (task *Task) RealUser() string {
	return GetUserNameById(task.User)
}

func AddTask(task *Task) error {
	ormer := orm.NewOrm()
	_, err := ormer.Insert(task)
	return err
}

func QueryTask(query string) []*Task {
	tasks := make([]*Task, 0)
	if query != "" {
		ormer := orm.NewOrm()
		qs := ormer.QueryTable(&Task{})
		cond := orm.NewCondition()
		cond = cond.Or("id__exact", query)
		cond = cond.Or("name__icontains", query)
		cond = cond.Or("status__exact", query)
		cond = cond.Or("content__icontains", query)
		qs.SetCond(cond).All(&tasks)
	}

	return tasks
}

func GetTaskById(id string) *Task {
	tid, _ := strconv.Atoi(id)
	task := &Task{ID: tid}
	ormer := orm.NewOrm()
	ormer.Read(task, "ID")
	return task
}

func UpdateTask(t *Task) {
	ormer := orm.NewOrm()
	ormer.Update(t)
}

func DeleteTask(id string) {
	tid, _ := strconv.Atoi(id)
	task := &Task{ID: tid}
	ormer := orm.NewOrm()
	ormer.Delete(task)
}
