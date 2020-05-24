package srv

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type JsonService struct {
	serialObj interface{}
}

func NewJsonService(serialObj interface{}) *JsonService {
	return &JsonService{serialObj: serialObj}
}

func (j *JsonService) Encode(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer f.Close()
	buffer := bufio.NewWriter(f)
	encoder := json.NewEncoder(buffer)
	err = encoder.Encode(j.serialObj)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = buffer.Flush()
	return nil
}
func (j *JsonService) Decode(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	buffer := bufio.NewReader(f)
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(j.serialObj)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonService) GetObj() interface{} {
	return j.serialObj
}

func (j *JsonService) String() string {
	t := reflect.TypeOf(j.serialObj)
	v := reflect.ValueOf(j.serialObj)
	return toString(t, v)
}


