package todolist

type Todo struct {
	Id        string
	Name      string
	StartTime string
	EndTime   string
	Status    string
	User      string
}

func NewTodo(id, name, startTime, endTime, status, user string) *Todo {
	return &Todo{
		Id:        id,
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
		Status:    status,
		User:      user,
	}
}

type TodoMgr struct {
	AllTodo []*Todo
}
