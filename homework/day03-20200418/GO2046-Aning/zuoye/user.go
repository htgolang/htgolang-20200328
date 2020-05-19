package main

import (
	"fmt"
	"strings"
)

func add(pk int, user map[string]map[string]string) {
	var (
		id string = fmt.Sprintf("%d", pk)
		//id string = strconv.Itoa(pk)
		name string
		age  string
		tel  string
		addr string
	)
	fmt.Println(id)
	fmt.Print("pkease input name: ")
	fmt.Scan(&name)
	fmt.Print("pkease input age: ")
	fmt.Scan(&age)
	fmt.Print("pkease input tel: ")
	fmt.Scan(&tel)
	fmt.Print("pkease input addr: ")
	fmt.Scan(&addr)

	user[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
	fmt.Println("add  success")
	fmt.Printf("%5s|%10s|%10s|%10s|%10s", "id", "name", "age", "addr", "tel\n")
	fmt.Printf("%5s|%10s|%10s|%10s|%10s", user[id]["id"], user[id]["name"], user[id]["age"], user[id]["tel"], user[id]["addr"])
}
func query(users map[string]map[string]string) {
	var q string
	fmt.Print("please input query: ")
	fmt.Scan(&q)
	for _, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["age"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["addr"], q) {
			fmt.Printf("%5s|%10s|%10s|%10s|%10s", "id", "name", "age", "addr", "tel\n")
			fmt.Printf("%5s|%10s|%10s|%10s|%10s", user["id"], user["name"], user["age"], user["addr"], user["tel"])
			fmt.Println()
		}
	}
}
func drop(user map[string]map[string]string) {

	var did, yes string

	fmt.Println("please input  [id|O]: ")
	fmt.Scan(&did)
	fmt.Println(user)

	if did == "O" {
		fmt.Println("delete all!!!")
		user = make(map[string]map[string]string)
		fmt.Println("all delete success!")
	} else if user1, ok := user[did]; ok {
		fmt.Println(user1["id"], user1["name"], user1["age"], user1["tel"], user1["addr"])
		fmt.Println("Did you del it?Y|N")
		fmt.Scan(&yes)
		if yes == "Y" || yes == "y" {
			delete(user, did)
			fmt.Println("del success")
		}
	} else {
		fmt.Println("the user[", did, "] not exits")
	}
}

func auth() bool {

	var password string = "centos"
	var passwd string
	var time int = 3

	//三次计数
	for i := 1; i <= time; i++ {
		fmt.Println("please input passwd:")
		fmt.Scan(&passwd)
		if password == passwd {
			return true
		} else {
			fmt.Println("the passwd is error!")
			fmt.Println("the times having ", time-i)
		}
		if i == time {
			fmt.Println("the error time more then three.")
			fmt.Println("exit")
			return false
		}
	}
	return false
}

func edit(user map[string]map[string]string) {
	var (
		id   string
		name string
		age  string
		tel  string
		addr string
		yes  string
	)
	fmt.Println("please input edit id :")
	fmt.Scan(&id)
	if user1, ok := user[id]; ok {
		fmt.Printf("%5s|%10s|%10s|%10s|%10s", "id", "name", "age", "addr", "tel\n")
		fmt.Printf("%5s|%10s|%10s|%10s|%10s\n", user1["id"], user1["name"], user1["age"], user1["addr"], user1["tel"])
		fmt.Println("edit it? Y|N")
		fmt.Scan(&yes)
		if yes == "Y" || yes == "y" {
			fmt.Print("pkease input name: ")
			fmt.Scan(&name)
			fmt.Print("pkease input age: ")
			fmt.Scan(&age)
			fmt.Print("pkease input tel: ")
			fmt.Scan(&tel)
			fmt.Print("pkease input addr: ")
			fmt.Scan(&addr)

			user[id] = map[string]string{
				"id":   id,
				"name": name,
				"age":  age,
				"tel":  tel,
				"addr": addr,
			}
			fmt.Println("edit success!")
		}
	} else {
		fmt.Println("id not exits")
	}
}
func main() {
	//认证
	if !auth() {
		fmt.Println("paswd error")
		return
	}
	fmt.Println("welcome to bb web")

	user := make(map[string]map[string]string)
	id := 0
	for {
		var op string
		fmt.Println(`	
		1.add							
		2.edit								
		3.del								
		4.select								
		5.exit
		please input opstion`)
		fmt.Scan(&op)
		if op == "1" {
			id++
			add(id, user)
		} else if op == "2" {
			edit(user)
		} else if op == "3" {
			drop(user)
		} else if op == "4" {
			query(user)
		} else if op == "5" {
			break
		} else {
			fmt.Println("commend is error")
		}

	}
}
