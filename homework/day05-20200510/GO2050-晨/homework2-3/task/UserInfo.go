package task

type User struct {
	name string
	addr string
	tel string
}

func NewUser(name,addr,tel string) *User  {
	return &User{
		name: name,
		addr: addr,
		tel: tel,
	}
}
func (user *User)SetUserTel(tel string)  {
	user.tel = tel
}
func (user *User)SetUserAddr(addr string)  {
	user.addr = addr
}
func (user *User)SetUserName(name string)  {
	user.name = name
}
func (user *User)GetUserTel() string  {
	return user.tel
}
func (user *User)GetUserAddr() string {
	return user.addr
}
func (user *User)GetUserName() string  {
	return user.name
}
