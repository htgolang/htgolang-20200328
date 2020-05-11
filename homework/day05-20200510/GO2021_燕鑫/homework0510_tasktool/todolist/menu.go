package todolist

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func TaskMain() {
	dburl := DBFILE
	user := login()
	if user == "" {
		return
	}
	tsrv := NewTaskService(dburl, user)
	for {
		if user == "root" {
			fmt.Println("0:ADMINUSERS", "1:SELECT", "2:UPDATE", "3:CREATE", "4:DELETE", "5:COMMIT", "6:RELOAD", "7:TASKCOUNT", "q:QUIT")
		} else {
			fmt.Println("1:SELECT", "2:UPDATE", "3:CREATE", "4:DELETE", "5:COMMIT", "6:RELOAD", "7:TASKCOUNT", "q:QUIT")
		}
		fmt.Print("Which action would you like to do: ")
		inputReader := bufio.NewReader(os.Stdin)
		action, _ := inputReader.ReadString('\n')
		switch strings.TrimSpace(action) {
		case "0":
			if user != "root" {
				fmt.Println("Only can enter 1,2,3,4,5,6,q")
			} else {
				adminUSERS()
			}
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
			fmt.Println("BYE!")
			return
		default:
			if user == "root" {
				fmt.Println("Only can enter 0,1,2,3,4,5,6,q")
			} else {
				fmt.Println("Only can enter 1,2,3,4,5,6,q")
			}
		}
	}
}

func adminUSERS() {
	accounturl := ACCOUNTFILE
	usersrv := NewUserService(accounturl)
	for {
		fmt.Println("1:SHOW USERS", "2:CREATE USER", "3:UPDATE USER", "4:DELETE USER", "q:QUIT")
		fmt.Print(`Please choose your "admin user" action: `)
		inputReader := bufio.NewReader(os.Stdin)
		action, _ := inputReader.ReadString('\n')
		switch strings.TrimSpace(action) {
		case "1":
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"id", "username", "createtime", "updatetime"})
			for _, user := range usersrv.users {
				table.Append([]string{user.Id, user.Username, user.CreateTime, user.UpdateTime})
			}
			table.Render()
		case "2":
			fmt.Print("ENTER NEW USERNAME: ")
			username, _ := inputReader.ReadString('\n')
			username = strings.TrimSpace(username)
			fmt.Print("ENTER YOUR PASSWORD: ")
			passwdbytes, _ := gopass.GetPasswd()
			password := string(passwdbytes)
			fmt.Print("ENTER YOUR PASSWORD AGAIN: ")
			passwdbytes, _ = gopass.GetPasswd()
			password2 := string(passwdbytes)
			if password != password2 {
				fmt.Println("The passwords entered are inconsistent")
				continue
			}
			err := usersrv.CreateUser(username, password)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "3":
			fmt.Println("UPDATE OK")
		case "4":
			fmt.Print("ENTER USERNAME OF WHICH YOU WANNA DELETE: ")
			username, _ := inputReader.ReadString('\n')
			username = strings.TrimSpace(username)
			err := usersrv.DeleteUser(username)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("USER %s has been  deleted successfully!\n", username)
		case "q":
			return
		default:
			fmt.Println("Only can enter 1,2,3,4,q")
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

		result, resultstr, err, sortkey, desc := tsrv.GetByFilter(filterstr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tsrv.printLines(result, sortkey, desc)
		fmt.Println(resultstr)
	}
}

func updateAction(tsrv *TaskService) {
	fmt.Println("Recent task list as following show:")
	result, resultstr, err, sortkey, desc := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result, sortkey, desc)
	fmt.Println(resultstr)
	fmt.Print("\nWhich task you wanna change,please enter the task id: ")
	inputReader := bufio.NewReader(os.Stdin)
	taskid, _ := inputReader.ReadString('\n')
	taskid = strings.TrimSpace(taskid)
	resultItem, _, err, sortkey, desc := tsrv.GetByFilter(fmt.Sprintf("id=%s", taskid))
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	status := ""
	user := ""
	fmt.Println("You can only change the status and users of the specific task: ")
	fmt.Print("status:")
	_, _ = fmt.Scanln(&status)
	fmt.Print("user:")
	_, _ = fmt.Scanln(&user)
	resultstr, err = tsrv.UpdateTask(resultItem[0], status, user)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	fmt.Println(resultstr)
	fmt.Println("This task as following show:")
	result, resultstr, err, sortkey, desc = tsrv.GetByFilter(fmt.Sprintf("id=%s", taskid))
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result, sortkey, desc)
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
	result, resultstr, err, sortkey, desc := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result, sortkey, desc)
	fmt.Println(resultstr)
}

func deleteAction(tsrv *TaskService) {
	fmt.Println("Recent task list as following show:")
	result, resultstr, err, sortkey, desc := tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result, sortkey, desc)
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
	result, resultstr, err, sortkey, desc = tsrv.GetByFilter("")
	if err != nil {
		fmt.Println("\n", err, "\n")
		return
	}
	tsrv.printLines(result, sortkey, desc)
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

func login() (username string) {
	accounturl := ACCOUNTFILE
	usersrv := NewUserService(accounturl)
	fmt.Println("1:SIGN IN   2:SIGN UP")
	fmt.Print("Your choose is: ")
	inputReader := bufio.NewReader(os.Stdin)
	action, _ := inputReader.ReadString('\n')
	switch strings.TrimSpace(action) {
	case "1":
		for i := 0; i < 3; i++ {
			fmt.Print("USERNAME: ")
			username, _ = inputReader.ReadString('\n')
			username = strings.TrimSpace(username)
			fmt.Print("PASSWORD: ")
			passwdbytes, _ := gopass.GetPasswd()
			password := string(passwdbytes)
			returnstr, err := usersrv.VerifyUser(username, password)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(returnstr)
			return username
		}
		fmt.Println("You have enter wrong 3 times!BYE!")
		return ""
	case "2":
		for i := 0; i < 3; i++ {
			fmt.Print("ENTER NEW USERNAME: ")
			username, _ = inputReader.ReadString('\n')
			username = strings.TrimSpace(username)
			fmt.Print("ENTER YOUR PASSWORD: ")
			passwdbytes, _ := gopass.GetPasswd()
			password := string(passwdbytes)
			fmt.Print("ENTER YOUR PASSWORD AGAIN: ")
			passwdbytes, _ = gopass.GetPasswd()
			password2 := string(passwdbytes)
			if password != password2 {
				fmt.Println("The passwords entered are inconsistent")
				continue
			}
			err := usersrv.CreateUser(username, password)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Welcome", username)
			return username
		}
		fmt.Println("You have enter wrong 3 times!BYE!")
		return ""
	default:
		fmt.Println("You can only choose SIGN IN or SIGN UP!")
		return ""
	}
}
