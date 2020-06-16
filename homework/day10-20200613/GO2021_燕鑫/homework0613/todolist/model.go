package todolist

type task struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
	Status    string `json:"status"`
	User      string `json:"user"`
}

type user struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	CreateTime string `json:"createtime"`
	UpdateTime string `json:"updatetime"`
}
