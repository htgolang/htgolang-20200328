package module

import (
	"github.com/astaxie/beego/orm"
	"math"
	"strconv"
	"time"
	"math/rand"
)

func salt() string{
	rand.Seed(time.Now().Unix())
	var salts string
	for i:= 0;i<10;i++{
		num := strconv.Itoa(rand.Intn(9))
		salts+=num
	}
	return  salts
}



type Manage struct {
	Id int     //id
	//Name string     `orm:"size(20)` //名字
	StartTime time.Time    `orm:"auto_now_add;type(datetime)"`//开始时间
	StopTime time.Time    `orm:"default(' ')"`//停止时间
	TaskName string       `orm:"size(20)"` //任务名称
	Taskinfo string     `orm:"size(100)"`   //任务描述
	TaskStatus int   `orm:"default(0)"`    //任务状态  0创建 1进行中 2暂停 3完成 4失败
	//default这里的默认值指的是当字段为null的时候默认为多少 如果字段默认是0那么default就不起到作用
	User   *User `orm:"rel(fk)"`
}

type UserMethod interface {
	GetUser(username string,field string)
	InputUser() (string,bool)
}

//根据权限返回权限对应的字符串
func Role(num int) string{
	switch num {
	case 0:
		return "普通用户"
	case 1:
		return "管理员"
	case 2:
		return "超级管理员"
	}
	return ""
}

//根据字符串返回对应的权限int
func GetRole(role string) int{
	switch role {
	case "普通用户":
		return 0
	case "管理员":
		return 1
	case "超级管理员":
		return 2
	}
	return 0
}


func Page(tablename string,index ,pagesize int) (int64,int,float64,int,error) {
	o := orm.NewOrm()
	//count 总共的数据量
	count ,err := o.QueryTable(tablename).Count()
	if err != nil {
		return 0,0,0,0,err
	}
	//起始量  因为从0开始所以-1  数据量起始点是0
	start := pagesize*(index -1 )
	//总共的页数= 总数据数量/ 单页数据量
	pagenum := math.Ceil(float64(count)/float64(int64(pagesize)))
	//总数据数量   起始点   总共的页数    每页数量
	return count,start,pagenum,pagesize,nil
}