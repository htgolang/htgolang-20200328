package forms

import "cmdb/models"

type NodeRegisterForm struct {
	UUID     string `json:"uuid"`
	Hostname string `json:"hostname"`
	Addr     string `json:"addr"`
}

type JobCreateForm struct {
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type JobModifyForm struct {
	ID     int    `form:"id"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type TargetCreateForm struct {
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}

type TargetModifyForm struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}

func NewTargetModifyForm(target *models.Target) *TargetModifyForm {
	form := &TargetModifyForm{}
	form.ID = target.ID
	form.Name = target.Name
	form.Remark = target.Remark
	form.Addr = target.Addr
	form.Job = target.Job.ID
	return form
}
