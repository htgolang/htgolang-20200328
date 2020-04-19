package question4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TaskMain() {
	dburl := DBFILE
	user := "yanxin"
	tsrv := NewTaskService(dburl, user)
	for {
		fmt.Println("1:SELECT", "2:UPDATE", "3:CREATE", "4:DELETE", "5:COMMIT", "6:RELOAD", "7:TASKCOUNT")
		fmt.Print("Which action would you like to do: ")
		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		action, _ := inputReader.ReadString('\n')
		switch strings.TrimSpace(action) {
		case "1":
			selectAction(tsrv)
		case "2":
			updateAction(tsrv)
		case "3":
			createAction(tsrv)
		case "4":
			deleteAction(tsrv)
		case "5":
			commitAction(tsrv)
		case "6":
			reloadAction(tsrv)
		case "7":
			showTaskCountAction(tsrv)
		case "q":
			return
		default:
			fmt.Println("Only can enter 1,2,3,4,5,6,q")
		}
	}
}

func selectAction(tsrv *TaskService) {
	for {
		fmt.Print("Please enter the filter condition(eg. id=1 name=task1 status=created): ")
		inputReader := bufio.NewReader(os.Stdin)
		filterstr, _ := inputReader.ReadString('\n')
		filterstr = strings.TrimSpace(filterstr)
		if filterstr == "q" {
			break
		}
		result, resultstr, err := tsrv.GetByFilter(filterstr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tsrv.printLines(result)
		fmt.Println(resultstr)
	}
}

func updateAction(tsrv *TaskService) {
	fmt.Println("Recent task list as following show:")
	result, resultstr, err := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result)
	fmt.Println(resultstr)
	fmt.Print("\nWhich task you wanna change,please enter the task id: ")
	inputReader := bufio.NewReader(os.Stdin)
	taskid, _ := inputReader.ReadString('\n')
	taskid = strings.TrimSpace(taskid)
	resultItem, _, err := tsrv.GetByFilter(fmt.Sprintf("id=%s", taskid))
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	status := ""
	user := ""
	fmt.Println("You can only change the status and user of the specific task: ")
	fmt.Print("status:")
	fmt.Scanln(&status)
	fmt.Print("user:")
	fmt.Scanln(&user)
	resultstr, err = tsrv.UpdateTask(resultItem[0], status, user)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println(resultstr)
	fmt.Println("This task as following show:")
	result, resultstr, err = tsrv.GetByFilter(fmt.Sprintf("id=%s", taskid))
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result)
	fmt.Println(resultstr)
}

func createAction(tsrv *TaskService) {
	fmt.Print("Enter a new task name: ")
	inputReader := bufio.NewReader(os.Stdin)
	taskName, _ := inputReader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	err := tsrv.CreateNewTask(taskName)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println("New task list as following show:")
	result, resultstr, err := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result)
	fmt.Println(resultstr)
}

func deleteAction(tsrv *TaskService) {
	fmt.Println("Recent task list as following show:")
	result, resultstr, err := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result)
	fmt.Println(resultstr)
	fmt.Print("\nWhich task you wanna change,please enter the task id: ")
	inputReader := bufio.NewReader(os.Stdin)
	taskid, _ := inputReader.ReadString('\n')
	taskid = strings.TrimSpace(taskid)
	resultstr, err = tsrv.DeleteTask(taskid)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println(resultstr)
	fmt.Println("New task list as following show:")
	result, resultstr, err = tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result)
	fmt.Println(resultstr)
}

func commitAction(tsrv *TaskService) {
	resultstr, err := tsrv.commit()
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println(resultstr)
}

func reloadAction(tsrv *TaskService) {
	resultstr, err := tsrv.reload()
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println(resultstr)
}

func showTaskCountAction(tsrv *TaskService) {
	fmt.Printf("\nThere are total %d tasks!\n\n", tsrv.GetTaskcount())
}
