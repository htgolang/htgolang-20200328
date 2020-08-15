package forms

import (
	"cmdb/models"
	"cmdb/utils"
	"encoding/json"
	"net/url"
	"time"
)

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

type AlertForm struct {
	Fingerprint  string            `json:"fingerprint"`
	Status       string            `json:"status"`
	StartsAt     *time.Time        `json:"startsAt"`
	EndsAt       *time.Time        `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
}

func (f *AlertForm) IsNew() bool {
	return f.Status == "firing"
}

func (f *AlertForm) AlertName() string {
	return f.Labels["alertname"]
}

func (f *AlertForm) LabelsString() string {
	if bytes, err := json.Marshal(f.Labels); err == nil {
		return string(bytes)
	}
	return "{}"
}

func (f *AlertForm) AnnotationsString() string {
	if bytes, err := json.Marshal(f.Annotations); err == nil {
		return string(bytes)
	}
	return "{}"
}

type AlertQueryParams struct {
	utils.PageQueryParams

	Q      string `form:"q"`
	Status string `form:"status"`
	Stime  string `form:"stime"`
	Etime  string `form:"etime"`
}

func NewAlertQueryParams(inputs url.Values) *AlertQueryParams {
	return &AlertQueryParams{PageQueryParams: utils.PageQueryParams{Inputs: inputs}}
}

func (f *AlertQueryParams) StartTime() *time.Time {
	loc, _ := time.LoadLocation("PRC") // 写入到配置文件
	if t, err := time.ParseInLocation("2006-01-02T15:04", f.Stime, loc); err == nil {
		// if t, err := time.Parse("2006-01-02T15:04", f.Stime); err == nil {
		return &t
	}
	return nil
}

func (f *AlertQueryParams) EndTime() *time.Time {
	loc, _ := time.LoadLocation("PRC")
	if t, err := time.ParseInLocation("2006-01-02T15:04", f.Etime, loc); err == nil {
		// if t, err := time.Parse("2006-01-02T15:04", f.Etime); err == nil {
		return &t
	}
	return nil
}
