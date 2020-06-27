package todolist

type task struct {
	Id        int64  `gorm:"column:Id;PRIMARY_KEY"`
	Name      string `gorm:"column:Name"`
	StartTime string `gorm:"column:StartTime"`
	EndTime   string `gorm:"column:EndTime"`
	Status    string `gorm:"column:Status"`
	User      string `gorm:"column:User"`
}

func (t *task) TableName() string {
	return "task"
}

type user struct {
	Id         int64  `gorm:"column:Id;PRIMARY_KEY"`
	Username   string `gorm:"column:Username"`
	Password   string `gorm:"column:Password"`
	Salt       string `gorm:"column:Salt"`
	CreateTime string `gorm:"column:CreateTime"`
	UpdateTime string `gorm:"column:UpdateTime"`
}

func (u *user) TableName() string {
	return "user"
}
