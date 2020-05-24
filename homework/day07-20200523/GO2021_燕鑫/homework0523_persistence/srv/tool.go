package srv

import (
	"fmt"
	"reflect"
	"strings"
)

func toString(t reflect.Type, v reflect.Value) string {
	switch t.Kind() {
	case reflect.Ptr:
		return toString(t.Elem(), v.Elem())
	case reflect.Struct:
		resultstr := make([]string, t.NumField()+2)
		resultstr[0] = "{"
		for i := 0; i < t.NumField(); i++ {
			resultstr[i+1] = fmt.Sprintf("%-20s:  %-20v", t.Field(i).Name, v.Field(i))
		}
		resultstr[len(resultstr)-1] = "}"
		return strings.Join(resultstr, "\n")
	case reflect.Slice:
		resultstr := []string{}
		for i:=0;i<v.Len();i++ {
			resultstr = append(resultstr, toString(t.Elem(), v.Index(i)))
		}
		return strings.Join(resultstr, "\n")
	default:
		return ""
	}
}