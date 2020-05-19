package task

type User struct {
	Name string
	Addr string
	Tel string
}

func NewUser(name,addr,tel string) *User  {
	return &User{
		Name: name,
		Addr: addr,
		Tel: tel,
	}
}
func (user *User)SetUserTel(tel string)  {
	user.Tel = tel
}
func (user *User)SetUserAddr(addr string)  {
	user.Addr = addr
}
func (user *User)SetUserName(name string)  {
	user.Name = name
}
func (user *User)GetUserTel() string  {
	return user.Tel
}
func (user *User)GetUserAddr() string {
	return user.Addr
}
func (user *User)GetUserName() string  {
	return user.Name
}
