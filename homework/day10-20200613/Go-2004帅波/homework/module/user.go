package module

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)




type User struct {
	Id int
	Name string  `orm:"size(20)";json:"name"`   //用户名
	Email  string  `orm:"size(100)";json:"email"` //邮箱 用来注册账号
	Password string `orm:"size(100)";json:"password"`   //密码
	Phone string  `orm:"size(100)";json:"phone"`
	Salt string   `orm:"size(100)";json:"salt"`//盐
	Role int   `orm:"default(0)"`
	Manage []*Manage `orm:"reverse(many)"`  //多任务
}


func (user User) TableName() string {
	return userTableName
}


func (user *User) GetId(mold interface{}) error{
	//获取到就返回user获取不到返回User的一个空值
	if mold.(*User) == nil {
		return errors.New("传入的是一个空值")
	}
	user = mold.(*User)

	if err = Ormer().Read(user);err != nil{
		return errors.New("查询失败")
	}
	return nil
}

func (user *User) Update(mold interface{}) error {
	var (
		email string
		phone string
		name string
	)
	if mold.(*User) == nil {
		return  errors.New("获取失败,传入的是空值")
	}
	user = mold.(*User)
	//这里没有做重复检测  随后会加
	email = user.Email
	phone = user.Phone
	name =  user.Name

	user.GetId(user)
	user.Email = email
	user.Name = name
	user.Phone = phone
	_,err = Ormer().Update(user)
	if err != nil {
		return errors.New("更新失败了")
	}
	return  nil
}


func (user User) Del(id int)  error{
	user.Id = id
	_,err = Ormer().Delete(&user)
	if err != nil {
		return  errors.New("删除失败")
	}
	return nil
}





func (user User) Add(mold interface{})  error {
	var (
		inputuser User
	)
	if mold.(*User) == nil {
		return  errors.New("传入的是空值")
	}
	user = *mold.(*User)
	o := orm.NewOrm()
	salts := salt()
	//判断用户信息是否重复
	inputuser.Name = user.Name
	if err = o.Read(&inputuser,"Name");err == nil {
		return errors.New("Name重复")
	}
	inputuser.Phone = user.Phone
	if err = o.Read(&inputuser,"Phone");err == nil {
		return errors.New("Phone重复")
	}
	inputuser.Email = user.Email
	if err = o.Read(&inputuser,"Email");err == nil {
		return errors.New("Email重复")
	}
	//加盐 并且把密码加密
	user.Salt = salts
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(salts+user.Password)))
	_,err = o.Insert(&user)
	if err != nil {
		return  err
	}
	return  nil
}

func (user User) ChangePass(id int,oldpass ,newpass string) error {
	user.Id = id
	if err = user.GetId(&user);err != nil {
		return errors.New("用户不存在")
	}
	//修改密码分为两种情况 一种是用户修改自己账号密码 那么需要旧密码跟新密码做判断
	//一种是管理员对用户密码初始化 这种情况进当作是 用户密码忘记 需要管理员帮忙重置密码
	if oldpass == " "{
		user.Salt = salt()

		user.Password =  fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+newpass)))
		if _,err = Ormer().Update(&user,"Password","Salt");err != nil{
			return errors.New("密码更新失败")
		}
	}else {
		oldpass = fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+oldpass)))
		if oldpass != user.Password {
			return errors.New("密码错误")
		}
		user.Salt = salt()
		user.Password =   fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+newpass)))
		if _,err = Ormer().Update(&user,"Password","Salt");err != nil{
			return errors.New("密码更新失败")
		}
	}
	return  nil
}

func (user *User) ComparePass(passwd string) error{
	err = Ormer().Read(user,"Email")
	if err != nil {
		return  errors.New("用户不存在")
	}
	//密码确认  read从数据库获取user的信息
	//用户输入的密码+ 数据库存储 的盐 然后md5  加密之后是否与数据库密码一致  一致返回true
	passwd = fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+passwd)))
	if passwd == user.Password {
		return nil
	}
	return errors.New("密码错误")
}


