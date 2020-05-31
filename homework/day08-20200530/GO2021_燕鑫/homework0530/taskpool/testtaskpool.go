package taskpool

import "fmt"

func TestTaskPool()  {
	taskpool:=NewTaskPool(10)
	createTask:= func(i int) func() interface{}{
		return func() interface{} {
			return fmt.Sprintf("func%d",i)
		}
	}

	for i:=0;i<5;i++{
		err:=taskpool.AddTask(createTask(i))
		if err!=nil{
			break
		}
	}
	err:=taskpool.AddTask(f1)
	if err!=nil{
		fmt.Println(err)
	}

	taskpool.Run()
	taskpool.GetResult()
}

func f1() interface{} {
	return "f1"
}