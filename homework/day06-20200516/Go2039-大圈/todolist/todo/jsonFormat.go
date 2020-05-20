package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

//以json格式保存数据到文件
func (msTodo *msTodo) jsonSaveToFile(f string) {
	//创建一个文件类型指针
	//打开文件，读写，追加，如果文件不存在则创建
	file, _ := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	//创建一个Encode类型指针，Encode对象
	jsonEncode := json.NewEncoder(file)

	//编码，Encode内存中的对象(msgTodo)到文件
	//jsonEncode.Encode(msgTodo.todoItems)
	jsonEncode.Encode(msTodo.todoItems)
}

//以json格式保存数据到文件
func (msTodo *msTodo) jsonReadToFile(f string) {
	//创建一个文件类型指针
	//打开文件，读写，追加，如果文件不存在则创建
	file, _ := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()


	//创建一个Decode类型指针，Decode对象
	jsonDecode := json.NewDecoder(file)

	//编码，Decode内存中的对象(msgTodo)到文件
	//循环遍历Decode对象，因为每次Decode时只有一行内容(代码得出的结果，是否正确待验证)
	for {
		err := jsonDecode.Decode(&msTodo.todoItems)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}