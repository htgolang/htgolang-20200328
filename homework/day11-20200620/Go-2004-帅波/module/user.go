package module

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)




type User struct {
	gorm.Model
	Name string  `gorm:"type:varchar(30);not null";json:"name"`   //用户名
	Email  string  `gorm:"type:varchar(30);not null";json:"email"` //邮箱 用来注册账号
	Password string `gorm:"type:varchar(40);not null";json:"password"`   //密码
	Phone string  `gorm:"type:varchar(30);not null";json:"phone"`
	Salt string   `gorm:"type:varchar(30);not null";json:"salt"`//盐
	Role int   `gorm:"default(0);not null"`
	//Manage []*Manage `orm:"reverse(many)"`  //多任务
}


func (user User) TableName() string {
	return userTableName
}



func (user *User) GetId(mold interface{}) error{
	//获取到就返回user获取不到返回User的一个空值
	if mold.(*User) == nil {
		return errors.New("传入的是一个空值")
	}
	//类型判断
	user,ok = mold.(*User)
	if !ok {
		return fmt.Errorf("类型错误不是user类型")
	}
	err = Db.Model(&User{}).Where("id = ? ",user.ID).First(&user).Error
	if err != nil {
		return fmt.Errorf("id 无效 %v\n",err)
	}
	return nil
}

func (user *User) Update(mold interface{}) error {
	var (
		phone ,email, name string
		users []User
	)
	user ,ok = mold.(*User)
	if !ok {
		return fmt.Errorf("类型错误不是user类型")
	}
	phone = user.Phone
	email = user.Email
	name = user.Name
	Db.Model(&User{}).Where("id != ? ",user.ID).Where(" name = ? or phone = ? or email = ?",name,phone,email).Find(&users)
	if len(users) > 0 {
		fmt.Println(users)
		return fmt.Errorf("信息重复")
	}
	//这里没有做重复检测  随后会加
	err = Db.Model(&User{}).First(&user,"id = ?",user.ID).Error
	if err != nil {
		return fmt.Errorf("user 赋值失败",err)
	}
	user.Email = email
	user.Phone = phone
	user.Name = name
	err = Db.Model(&User{}).Update(&user).Error
	if err != nil {
		return fmt.Errorf("user更新失败",err)
	}
	return  nil
}


func (user User) Del(id uint)  error{
	err = Db.Model(&User{}).Where("id = ? ",id).Delete(&User{}).Error
	if err != nil {
		return  errors.New("删除失败")
	}
	return nil
}

//约定 就是如果这个结构体有一个方法需要指针接受者 那么其他所有方法最好都是指针接受者方法
func (user *User) Add(mold interface{})  error {
	var (
		users []User
	)
	user ,ok = mold.(*User)
	if !ok {
		return fmt.Errorf("类型错误不是user类型")
	}
	salts := salt()
	//判断用户信息是否重复
	Db.Where("name = ? or email = ? or phone = ? ",user.Name,user.Email,user.Phone).Find(&users)
	if len(users) > 0  {
		return fmt.Errorf("信息重复 ")
	}
	//加盐 并且把密码加密
	user.Salt = salts
	user.Password = fmt.Sprintf("%x",md5.Sum([]byte(salts+user.Password)))
	err = Db.Create(&user).Error
	if err != nil {
		return  fmt.Errorf("user创建失败")
	}
	return  nil
}

func (user User) ChangePass(id uint,oldpass ,newpass string) error {
	user.ID = id
	if err = user.GetId(&user);err != nil {
		return fmt.Errorf("%v\n",err)
	}
	//修改密码分为两种情况 一种是用户修改自己账号密码 那么需要旧密码跟新密码做判断
	//一种是管理员对用户密码初始化 这种情况进当作是 用户密码忘记 需要管理员帮忙重置密码
	if oldpass == " "{
		user.Salt = salt()

		user.Password =  fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+newpass)))
		err = Db.Model(User{}).Update(&user).Error
		if  err != nil{
			return errors.New("密码更新失败")
		}
	}else {
		oldpass = fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+oldpass)))
		if oldpass != user.Password {
			return errors.New("密码错误")
		}
		user.Salt = salt()
		user.Password =   fmt.Sprintf("%x",md5.Sum([]byte(user.Salt+newpass)))
		if err = Db.Model(User{}).Update(&user).Error;err != nil{
			return errors.New("密码更新失败")
		}
	}
	return  nil
}

func (user *User) ComparePass(passwd string) error{
	Db.Where("email = ?",user.Email).First(&user)
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


func (user *User) Get(mold,value interface{}) error{
	err = Db.Model(&User{}).Where(" email = ? ",value).First(&user).Error
	if err != nil {
		return fmt.Errorf("获取失败",err)
	}
	return nil
}