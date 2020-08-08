package services

import (
	"cmdb/forms"
	"cmdb/models"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type nodeService struct {
}

func (s *nodeService) Register(form *forms.NodeRegisterForm) *models.Node {
	node := &models.Node{UUID: form.UUID}
	ormer := orm.NewOrm()
	if err := ormer.Read(node, "UUID"); err == nil {
		// 有数据更新
		node.Hostname = form.Hostname
		node.Addr = form.Addr
		node.DeletedAt = nil
		ormer.Update(node)
	} else if err == orm.ErrNoRows {
		//无数据创建
		node.Hostname = form.Hostname
		node.Addr = form.Addr
		ormer.Insert(node)
	} else {
		return nil
	}

	return node
}

// Query 查询
func (s *nodeService) Query(q string) []*models.Node {
	var nodes []*models.Node
	queryset := orm.NewOrm().QueryTable(&models.Node{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("hostname__icontains", q)
		qcond = qcond.Or("addr__icontains", q)
		cond = cond.AndCond(qcond)
	}
	queryset.SetCond(cond).All(&nodes)
	return nodes
}

func (s *nodeService) GetByPk(pk int) *models.Node {
	node := &models.Node{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(node); err == nil {
		return node
	}
	return nil
}

func (s *nodeService) Delete(pk int) {
	if node := s.GetByPk(pk); node != nil {
		now := time.Now()
		node.DeletedAt = &now
		orm.NewOrm().Update(node, "DeletedAt")
	}
}

type jobService struct {
}

// Query 查询
func (s *jobService) Query(q string) []*models.Job {
	var jobs []*models.Job
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Job{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("key__icontains", q)
		qcond = qcond.Or("remark__icontains", q)
		qcond = qcond.Or("node__hostname__icontains", q)
		qcond = qcond.Or("node__addr__icontains", q)
		cond = cond.AndCond(qcond)
	}
	queryset.RelatedSel().SetCond(cond).All(&jobs)
	return jobs
}

func (s *jobService) GetByPk(pk int) *models.Job {
	job := &models.Job{ID: pk}
	ormer := orm.NewOrm()

	if err := ormer.Read(job); err == nil {
		ormer.LoadRelated(job, "Node") // job.Node = NodeService.GetByPk(job.Node.ID)
		return job
	}
	return nil
}

func (s *jobService) Delete(pk int) {
	if job := s.GetByPk(pk); job != nil {
		now := time.Now()
		job.DeletedAt = &now
		orm.NewOrm().Update(job, "DeletedAt")
	}
}

func (s *jobService) Create(form *forms.JobCreateForm) *models.Job {
	job := &models.Job{
		Key:    form.Key,
		Remark: form.Remark,
		Node:   NodeService.GetByPk(form.Node),
	}
	if _, err := orm.NewOrm().Insert(job); err == nil {
		return job
	}
	return nil
}

func (s *jobService) Modify(form *forms.JobModifyForm) *models.Job {
	if job := s.GetByPk(form.ID); job != nil {
		job.Key = form.Key
		job.Remark = form.Remark
		job.Node = NodeService.GetByPk(form.Node)
		orm.NewOrm().Update(job)
		return job
	}
	return nil
}

func (s *jobService) QueryByUUID(uuid string) []*models.Job {
	var jobs []*models.Job
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.Job{})
	queryset.Filter("deleted_at__isnull", true).Filter("node__uuid", uuid).All(&jobs)
	for _, job := range jobs {
		ormer.LoadRelated(job, "Targets")
	}
	return jobs
}

type targetService struct {
}

// Query 查询
func (s *targetService) Query(q string) []*models.Target {
	var targets []*models.Target
	queryset := orm.NewOrm().QueryTable(&models.Target{})
	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)
	if q != "" {
		qcond := orm.NewCondition()
		qcond = qcond.Or("name__icontains", q)
		qcond = qcond.Or("remark__icontains", q)
		qcond = qcond.Or("addr__icontains", q)
		qcond = qcond.Or("job__key__icontains", q)
		qcond = qcond.Or("job__remark__icontains", q)
		qcond = qcond.Or("job__node__hostname__icontains", q)
		qcond = qcond.Or("job__node__addr__icontains", q)
		cond = cond.AndCond(qcond)
	}
	queryset.RelatedSel().SetCond(cond).All(&targets)
	return targets
}

func (s *targetService) GetByPk(pk int) *models.Target {
	target := &models.Target{ID: pk}
	if err := orm.NewOrm().Read(target); err == nil {
		return target
	}
	return nil
}

func (s *targetService) Delete(pk int) {
	if target := s.GetByPk(pk); target != nil {
		now := time.Now()
		target.DeletedAt = &now
		orm.NewOrm().Update(target, "DeletedAt")
	}
}

func (s *targetService) Create(form *forms.TargetCreateForm) *models.Target {
	target := &models.Target{
		Name:   form.Name,
		Addr:   form.Addr,
		Remark: form.Remark,
		Job:    JobService.GetByPk(form.Job),
	}
	if _, err := orm.NewOrm().Insert(target); err == nil {
		return target
	} else {
		fmt.Println(err, form)
	}
	return nil
}

func (s *targetService) Modify(form *forms.TargetModifyForm) *models.Target {
	if target := s.GetByPk(form.ID); target != nil {
		target.Name = form.Name
		target.Addr = form.Addr
		target.Remark = form.Remark
		target.Job = JobService.GetByPk(form.Job)
		orm.NewOrm().Update(target)
		return target
	}
	return nil
}

var (
	NodeService   = new(nodeService)
	JobService    = new(jobService)
	TargetService = new(targetService)
)
