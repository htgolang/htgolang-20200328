package module

type Operation interface {
	CRUD
	//返回一个表名  如果做关联查询好用
	TableName() string
	UserShare
}




type CRUD interface {
	Add(mold interface{}) error
	GetId(mold interface{})  error
	Update(mold interface{}) error
	Del(id int) error
}


type UserShare interface {
	ComparePass(passwd string) error
	ChangePass(id int,oldpass ,newpass string) error
}