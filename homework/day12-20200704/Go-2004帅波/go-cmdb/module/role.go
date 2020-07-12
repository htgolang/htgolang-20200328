package module


//权限这里有
/*
1.游客         登录成功之后所有不可看不可改 不可动
2.用户管理员      对用户的删改密码初始化操作
3.成员   		对自己密码 修改 以及修改个人基本信息
4.项目负责人    某一个项目的总负责人 可以对项目下面涉及到的所有资源进行操作
5.项目运维      类似于
6.项目开发     只允许对项目测试环境做发布
7.所有只读    所有页面可读 但是不可操作
8.超管        只有一个账号是超级管理员
*/

//这里默认是有12378
//至于项目这里的话需要到k8s这里取做
//基于ns做控制  一个项目一个ns所有只读  但是操作需要被授权

type Group struct {
	Id int     `json:"id"`
	Type int   `json:"type"`
	Comment string `orm:"size(20)";json:"comment"`
	//NameSpace string  `orm:"size(20)";json:"name_space"`
	User *User  `orm:"rel(fk)";json:"user"` //0普通用户  1管理员  2超管
}


func (group Group) TableName() string{
	return groupTableName
}










