package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Node struct {
	ID        int        `orm:"column(id);"`
	UUID      string     `orm:"column(uuid);varchar(64)"`
	Hostname  string     `orm:"varchar(64)"`
	Addr      string     `orm:"varchar(512)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`

	Jobs []*Job `orm:"reverse(many);"`
}

type Job struct {
	ID        int        `orm:"column(id);" json:"id"`
	Key       string     `orm:"varchar(64)" json:"key"`
	Remark    string     `orm:"varchar(512)" json:"remark"`
	CreatedAt *time.Time `orm:"auto_now_add" json:"-"`
	UpdatedAt *time.Time `orm:"auto_now" json:"-"`
	DeletedAt *time.Time `orm:"null" json:"-"`

	Node    *Node     `orm:"rel(fk)" json:"-"`
	Targets []*Target `orm:"reverse(many)" json:"targets"`
}

type Target struct {
	ID        int        `orm:"column(id);" json:"id"`
	Name      string     `orm:"varchar(64)" json:"name"`
	Remark    string     `orm:"varchar(512)" json:"remark"`
	Addr      string     `orm:"varchar(126)" json:"addr"`
	CreatedAt *time.Time `orm:"auto_now_add" json:"-"`
	UpdatedAt *time.Time `orm:"auto_now" json:"-"`
	DeletedAt *time.Time `orm:"null" json:"-"`

	Job *Job `orm:"rel(fk)" json:"-"`
}

func init() {
	orm.RegisterModel(new(Node), new(Job), new(Target))
}
