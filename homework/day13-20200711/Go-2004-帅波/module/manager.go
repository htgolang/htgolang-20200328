package module

import "fmt"

type Manager struct {
	//gorm.Model
	Name string
	//UserName string
}

var _ Operation = (*Manager)(nil)

func (m *Manager) TableName() string {
	return managerTableName
}


func (m *Manager) Add(mold interface{})  error{
	return  nil
}

func (m *Manager) GetId(mold interface{}) error{
	return nil
}

func (m *Manager) Update(mold interface{}) error {
	return nil
}

func (m *Manager) Del(id uint) error {
	return nil
}

func (m *Manager) Get(mold string,value interface{}) error {
	return nil
}

func (m *Manager) GetAll(mold interface{}) (error) {
	var managers *[]Manager
	managers = mold.(*[]Manager)
	jobs,err := jenkinsclient.GetAllJobs()
	//for _,v := range  jobs {
	//	fmt.Printf("%#v\n",v.Raw)
	//}
	fmt.Printf("%#v\n",jobs[0].Raw)
	if err != nil {
		return err
	}
	for _,v := range  jobs {
		m.Name = v.Raw.Name
		*managers = append(*managers,*m)
	}
	return  nil
}

func (m *Manager) UpdateMold(value interface{}) error {
	return nil
}

func (m *Manager) Query(method string,mold interface{}) error {
	return  nil
}



