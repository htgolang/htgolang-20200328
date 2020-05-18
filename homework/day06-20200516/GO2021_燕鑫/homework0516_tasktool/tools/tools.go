package tools

import (
	"fmt"
	"reflect"
	"strings"
)

func Contain(value string, obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		f := fmt.Sprintf("%s", t.Field(i).Name)
		if strings.ToLower(f) == strings.ToLower(value) {
			return fmt.Sprintf("%s", v.Field(i))
		}
	}
	return ""
}

func In(key string, sli []string) bool {
	for _, value := range sli {
		if value == key {
			return true
		}
	}
	return false
}
