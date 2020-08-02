package services

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"

	"cmdb/forms"
	"cmdb/models"
	"cmdb/utils"
)

type taskService struct {
}

//Add 添加任务
func (s *taskService) Add(form *forms.FormTask) error {
	task := &models.Task{ID: form.ID}
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

// Query 查询任务
func (s *taskService) Query(query string, user *models.User) []*models.Task {
	var tasks []*models.Task
	ormer := orm.NewOrm()
	qs := ormer.QueryTable(&models.Task{})
	cond := orm.NewCondition()

	if user.IsAdmin == 0 {
		cond = cond.And("user__iexact", user.ID)
	}

	if query != "" {
		cond = cond.Or("id__exact", query)
		cond = cond.Or("name__icontains", query)
		cond = cond.Or("status__exact", query)
		cond = cond.Or("content__icontains", query)
	}
	qs.SetCond(cond).All(&tasks)

	return tasks
}

func (s *taskService) GetById(id int) *models.Task {
	task := &models.Task{ID: id}
	ormer := orm.NewOrm()
	if ormer.Read(task, "ID") == nil {
		return task
	}
	return nil
}

func (s *taskService) Update(form *forms.FormTask) error {
	if task := s.GetById(form.ID); task != nil {
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

func (s *taskService) Delete(id int) error {
	task := &models.Task{ID: id}
	ormer := orm.NewOrm()
	_, err := ormer.Delete(task)
	return err
}

var TaskService = &taskService{}
