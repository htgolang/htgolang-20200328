package srv

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"reflect"
)

type GobService struct {
	serialObj interface{}
}

func NewGobService(serialObj interface{}) *GobService {
	return &GobService{serialObj: serialObj}
}


func(g *GobService) Encode(filepath string) error {
	f,err:=os.OpenFile(filepath,os.O_CREATE|os.O_TRUNC,os.ModePerm)
	defer f.Close()
	buffer:=bufio.NewWriter(f)
	encoder:=gob.NewEncoder(buffer)
	err=encoder.Encode(g.serialObj)
	if err!=nil {
		fmt.Println(err)
		return err
	}
	_=buffer.Flush()
	return nil
}
func (g *GobService) Decode(filepath string) (error){
	f,err:=os.Open(filepath)
	if err!=nil{
		return err
	}
	defer f.Close()
	buffer:=bufio.NewReader(f)
	decoder:=gob.NewDecoder(buffer)
	err=decoder.Decode(g.serialObj)
	if err!=nil{
		return err
	}
	return nil
}

func (g *GobService) GetObj() interface{} {
	return g.serialObj
}

func (g *GobService) String() string {
	t := reflect.TypeOf(g.serialObj)
	v := reflect.ValueOf(g.serialObj)
	return toString(t, v)
}
