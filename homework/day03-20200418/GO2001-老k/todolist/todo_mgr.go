package todolist

import "fmt"

func (t *TodoMgr) AddTodo(todo *Todo) {
	t.AllTodo = append(t.AllTodo, todo)
}

func (t *TodoMgr) EditTodolist(todo *Todo) {
	for i, v := range t.AllTodo {
		if v.Id == todo.Id {
			t.AllTodo[i] = todo
			fmt.Println("修改成功")
			return
		}
	}
	fmt.Println("没有找到%", todo.Id)
}

func (t *TodoMgr) ShowTodolist() {
	for _, v := range t.AllTodo {
		fmt.Printf("id:%s--name:%s--startTime:%s--endTime:%s--status:%s--user:%s\n",
			v.Id, v.Name, v.StartTime, v.EndTime, v.Status, v.User)
	}
}

func (t *TodoMgr) DeleteTodolist(todo *Todo) {
	for i, v := range t.AllTodo {
		if v.Id == todo.Id {
			t.AllTodo = append(t.AllTodo[:i], t.AllTodo[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Printf("没有找到%s这个任务", todo.Id)
}
