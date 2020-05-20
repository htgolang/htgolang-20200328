package todo

import (
	"encoding/gob"
	"fmt"
	"os"
)
//注册持久化的对象到gob管理器中
func init() {
	gob.Register(&msTodo{})
}
//保存数据到文件，以gob格式保存
func (msTodo *msTodo) gobSaveToFile(f string)  {
	//打开一个文件
	//打开文件，读写，追加，如果文件不存在则创建
	file, _ := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	//创建一个encode对象
	gobEncode := gob.NewEncoder(file)
	//编码内存中的数据源到文件中
	gobEncode.Encode(msTodo.todoItems)

}
//从文件中读取gob格式数据到msTodo中
func (msTodo *msTodo) gobReadToFile(f string) {
	//打开一个文件，返回文件类型指针
	//打开文件，读写，追加，如果文件不存在则创建
	file, _ := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	//创建一个Decode对象
	gobDecode := gob.NewDecoder(file)
	//解码
	gobDecode.Decode(&msTodo.todoItems)
	fmt.Println(msTodo.todoItems)
}
