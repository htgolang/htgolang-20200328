package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// TaskStatusRelation 任务状态关系类型
type TaskStatusRelation struct {
	Name     string
	Relation []int
}

// TaskStatusMap 任务状态map
var TaskStatusMap = map[int]TaskStatusRelation{
	1: TaskStatusRelation{Name: "新建", Relation: []int{1, 2}},
	2: TaskStatusRelation{Name: "正在进行", Relation: []int{2, 3, 4}},
	3: TaskStatusRelation{Name: "暂停", Relation: []int{2, 3, 4}},
	4: TaskStatusRelation{Name: "已完成", Relation: []int{4}},
}

func init() {
	orm.RegisterModel(new(Task))
}

// Task 任务对象
type Task struct {
	ID           int        `orm:"column(id);pk;auto"`
	Name         string     `orm:"column(name);description(任务名)"`
	StatusID     int        `orm:"column(status_id);description(任务状态ID)"`
	StartTime    *time.Time `orm:"column(start_time);description(开始时间)"`
	CompleteTime *time.Time `orm:"column(complete_time);type(datetime);description(完成时间)"`
	DeadlineTime *time.Time `orm:"column(Deadline_time);type(datetime);description(截止时间)"`
	UserID       int        `orm:"column(user_id);description(负责人ID)"`
	Describe     string     `orm:"column(describe);size(1024);description(任务描述)"`
}

// TableName 设置表名
func (t *Task) TableName() string {
	return "tasks"
}

// UpdateTask 更新用户信息
func (t Task) UpdateTask() error {
	var (
		err error
	)
	_, err = orm.NewOrm().Update(&t)
	return err
}

// DeleteTask 删除用户
func (t Task) DeleteTask() error {
	var (
		err error
	)
	_, err = orm.NewOrm().Delete(&t, "id")
	return err
}

// GetTaskByID 根据任务ID获取任务信息
func (t *Task) GetTaskByID() error {
	return orm.NewOrm().Read(t, "id")
}

// QueryTasks 可以根据关键字查询信息
func QueryTasks(key string) []Task {
	var (
		tasks    []Task
		queryset orm.QuerySeter
		cond     *orm.Condition
	)
	queryset = orm.NewOrm().QueryTable(&Task{})
	if key != "" {
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", key)
		cond = cond.Or("describe__icontains", key)
		queryset.SetCond(cond)
	}
	queryset.All(&tasks)
	return tasks
}
