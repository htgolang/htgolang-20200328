package load_save

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"todolist/operations"
	"todolist/task"
)

func Load_option() []*task.Task {
	var f Factory
	var text string = operations.Input("请输入读取方式(json/gob)")
	var opt string = "tasks"
	if text == "gob"{
		f = &GobFactory{}
		opt = opt + ".gob"
	}else if text =="json"{
		f = &JsonFactory{}
		opt = opt + ".json"
	} else {
		fmt.Println("Invalid selection")
		return nil
	}
	t := f.Create()
	return t.Load(opt)
}

func Save_option(tasks []*task.Task)  {
	var text string = operations.Input("请输入保存方式(json/gob)")
	var f Factory
	if text == "gob"{
		f = &GobFactory{}
	} else if text == "json"{
		f = &JsonFactory{}
	} else {
		fmt.Println("Invalid save option!")
	}
	t := f.Create()
	t.Save(tasks)
}

type SaveLoad interface {
	Load(string) []*task.Task
	Save([]*task.Task)
}

type Json struct {

}

func (j Json)Load(path string) []*task.Task   {
	var tasks []*task.Task
	jsonTxt,_ := ioutil.ReadFile(path)

	json.Unmarshal(jsonTxt,&tasks)

	return tasks
}
func (j Json)Save(tasks []*task.Task)  {
	ctx,_ := json.Marshal(tasks)
	var buffer bytes.Buffer
	json.Indent(&buffer,ctx,"","\t")
	file,_ := os.Create("tasks.json")
	defer file.Close()
	buffer.WriteTo(file)
}

type Gob struct {

}

func (g Gob)Load(path string) []*task.Task {
	var tasks []*task.Task
	file,_ :=os.Open(path)
	defer file.Close()

	decoder := gob.NewDecoder(file)
	decoder.Decode(&tasks)

	for _,task :=range tasks {
		fmt.Println(task)
	}
	return tasks
}
func (g Gob)Save(tasks []*task.Task)  {
	file,err := os.Create("tasks.gob")
	if err != nil {
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(tasks)

}

type Factory interface {
	Create() SaveLoad
}

type GobFactory struct {
}

func (g *GobFactory)Create()  SaveLoad {
	return Gob{}
}
type JsonFactory struct {
}

func (j *JsonFactory)Create() SaveLoad {
	return Json{}
}


