package module



type Operation interface {
	CRUD
}




type CRUD interface {
	Add(mold interface{}) error
	GetId(mold interface{})  error
	Update(mold interface{}) error
	Del(id int) error
}

type UserShare interface {
	ChangePass(id int,oldpass ,newpass string) error
}