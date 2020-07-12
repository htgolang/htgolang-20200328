package module

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name string  `gorm:"type:varchar(30);not null";json:"name"`   //用户名
	Email  string  `gorm:"type:varchar(30);not null";json:"email"` //邮箱 用来注册账号
	Password string `gorm:"type:varchar(100);not null";json:"password"`   //密码
	Phone string  `gorm:"type:varchar(30);not null";json:"phone"`
	//Salt string   `gorm:"type:varchar(30);not null";json:"salt"`//盐
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
	err = db.Model(&User{}).Where("id = ? ",user.ID).First(&user).Error
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
	db.Model(&User{}).Where("id != ? ",user.ID).Where(" name = ? or phone = ? or email = ?",name,phone,email).Find(&users)
	if len(users) > 0 {
		fmt.Println(users)
		return fmt.Errorf("信息重复")
	}
	//这里没有做重复检测  随后会加
	err = db.Model(&User{}).First(&user,"id = ?",user.ID).Error
	if err != nil {
		return fmt.Errorf("user 赋值失败",err)
	}
	user.Email = email
	user.Phone = phone
	user.Name = name
	err = db.Model(&User{}).Update(&user).Error
	if err != nil {
		return fmt.Errorf("user更新失败",err)
	}
	return  nil
}


func (user User) Del(id uint)  error{
	err = db.Model(&User{}).Where("id = ? ",id).Delete(&User{}).Error
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
	//判断用户信息是否重复
	db.Where("name = ? or email = ? or phone = ? ",user.Name,user.Email,user.Phone).Find(&users)
	if len(users) > 0  {
		return fmt.Errorf("信息重复 ")
	}
	//加盐 并且把密码加密
	user.Password ,err = HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("密码修改失败 %v\n",err)
	}
	err = db.Create(&user).Error
	if err != nil {
		return  fmt.Errorf("user创建失败 %v",err)
	}
	return  nil
}

func (user *User) Get(mold string,value interface{}) error {
	user ,ok = value.(*User)
	if !ok {
		return fmt.Errorf("类型错误不是user类型")
	}
	switch mold {
	case "email":
		err = db.Model(&User{}).Where(" email = ? ",user.Email).First(&user).Error
		if err != nil {
			return fmt.Errorf("获取失败%v\n",err)
		}
	}
	return nil
}


func (user *User) GetAll(mold interface{}) (interface{},error) {
	var users []User
	users ,ok = mold.([]User)
	if !ok {
		return nil,fmt.Errorf("类型错误不是users类型")
	}
	err = db.Model(&User{}).Find(&users).Error
	if err != nil {
		return nil,fmt.Errorf("查询失败 %v",err)
	}
	return users,nil
}

func (user *User) UpdateMold(value interface{}) error{
	user,ok  = value.(*User)
	if !ok {
		return fmt.Errorf("类型错误不是users类型")
	}
	err = db.Model(&User{}).Where("id = ?",user.ID).Update("password",user.Password).Error
	if err != nil {
		return fmt.Errorf("更新失败",err)
	}
	return  nil
}

//返回加密的密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
//密码验证 修改密码
func ChangePasswordHash(oldpass ,newpass string,user *User) error {
	operation :=  NewOperation(&User{})
	err := operation.Get("email",user)
	if err != nil {
		return fmt.Errorf("获取用户失败 %v\n",err)
	}
	if oldpass == "" {
		user.Password ,err = HashPassword(newpass)
		if err != nil {
			return fmt.Errorf("获取新密码失败 %v\n",err)
		}
		err = operation.UpdateMold(user)
		if err != nil {
			return fmt.Errorf("更新失败 %v\n",err)
		}
	}else {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldpass))
		if err != nil {
			return fmt.Errorf("密码错误 %v\n",err)
		}
		user.Password ,err = HashPassword(newpass)
		if err != nil {
			return fmt.Errorf("获取新密码失败 %v\n",err)
		}
		err = operation.UpdateMold(user)
		if err != nil {
			return fmt.Errorf("更新失败 %v\n",err)
		}
	}
	return nil
}

//密码验证
func CheckPassword(password string,user *User) error{
	operation :=  NewOperation(&User{})
	err := operation.Get("email",user)
	if err != nil {
		return  fmt.Errorf("获取用户失败 %v\n",err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("密码错误 %v\n",err)
	}
	return nil
}
