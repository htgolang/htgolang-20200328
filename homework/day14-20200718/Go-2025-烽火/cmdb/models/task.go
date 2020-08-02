package models

import (
	"errors"
	"time"

	"cmdb/forms"
	"cmdb/utils"

	"github.com/astaxie/beego/orm"
)

type Task struct {
	ID           int        `orm:"column(id)"`
	Name         string     `orm:"size(64)"`
	StartTime    *time.Time `orm:"column(start_time)"`
	CompleteTime *time.Time `orm:"column(complete_time);null"`
	DeadlineTime *time.Time `orm:"column(deadline_time);"`
	Status       int        `orm:""`
	Content      string     `orm:"null"`
	User         int        `orm:""`
	DeletedAt    *time.Time `orm:"column(deleted_at);null"`
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

func AddTask(form *forms.FormTask) error {
	task := &Task{ID: form.ID}
	task.Name = form.Name
	task.Status = form.Status
	task.Content = form.Content
	task.User = form.User
	task.StartTime = utils.String2Time(form.StartTime)
	task.DeadlineTime = utils.String2Time(form.DeadlineTime)
	task.CompleteTime = utils.String2Time(form.CompleteTime)

	ormer := orm.NewOrm()
	_, err := ormer.Insert(task)
	return err
}

func QueryTask(query string) []*Task {
	var tasks []*Task
	ormer := orm.NewOrm()
	qs := ormer.QueryTable(&Task{})

	if query != "" {
		cond := orm.NewCondition()
		cond = cond.Or("id__exact", query)
		cond = cond.Or("name__icontains", query)
		cond = cond.Or("status__exact", query)
		cond = cond.Or("content__icontains", query)
		qs.SetCond(cond).All(&tasks)
	} else {
		qs.All(&tasks)
	}

	return tasks
}

func GetTaskById(id int) *Task {
	task := &Task{ID: id}
	ormer := orm.NewOrm()
	if ormer.Read(task, "ID") == nil {
		return task
	}
	return nil
}

func UpdateTask(form *forms.FormTask) error {
	if task := GetTaskById(form.ID); task != nil {
		task.Name = form.Name
		task.Status = form.Status
		task.Content = form.Content
		task.User = form.User
		task.StartTime = utils.String2Time(form.StartTime)
		task.DeadlineTime = utils.String2Time(form.DeadlineTime)
		if form.Status == 3 {
			now := time.Now()
			task.CompleteTime = &now
		}
		ormer := orm.NewOrm()
		_, err := ormer.Update(task)
		return err
	}

	return errors.New("任务不存在")
}

func DeleteTask(id int) error {
	task := &Task{ID: id}
	ormer := orm.NewOrm()
	_, err := ormer.Delete(task)
	return err
}
