package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	id   int
	name string
}

func (u User) GetId() int {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

// 获取变量数据类型
func GetType(typ reflect.Type) string {
	// relfect.Type
	switch typ.Kind() {
	case reflect.Int, reflect.String, reflect.Float64, reflect.Bool:
		return typ.Name()
	case reflect.Array:
		//[3]int 元素的数据类型 typ.Elem() 基本数据类型直接显示名称
		return fmt.Sprintf("[%d]%s", typ.Len(), GetType(typ.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", GetType(typ.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s][%s]", GetType(typ.Key()), GetType(typ.Elem()))
	case reflect.Struct:
		/* type xxxx struct
		Fields:
			name, type tag
		Methods:
		*/
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("type %s struct\n", typ.Name()))
		builder.WriteString(fmt.Sprintf("\tFields(%d):\n", typ.NumField()))
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			name := field.Name + " "
			if field.Anonymous {
				name = ""
			}
			tag := ""
			if field.Tag != "" {
				tag = fmt.Sprintf(" `%s`", field.Tag)
			}
			builder.WriteString(fmt.Sprintf("\t\t%s%s%s\n", name, GetType(field.Type), tag))
		}
		builder.WriteString(fmt.Sprintf("\tMethods(%d):\n", typ.NumMethod()))
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			builder.WriteString(fmt.Sprintf("\t\tfunc %s%s\n", method.Name, "()"))
		}

		return builder.String()
	case reflect.Func:
		// (1,1,1,1, ...)(1)
		var builder strings.Builder
		builder.WriteString("(")
		for i := 0; i < typ.NumIn(); i++ {
			if i != 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(GetType(typ.In(i)))
		}
		if typ.IsVariadic() {
			if typ.NumIn() > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString("...")
		}
		builder.WriteString(")")
		if typ.NumOut() > 0 {
			builder.WriteString(" (")
			for i := 0; i < typ.NumOut(); i++ {
				if i != 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(GetType(typ.Out(i)))
			}
			builder.WriteString(")")
		}
		return builder.String()
	case reflect.Ptr:
		var builder strings.Builder
		builder.WriteString("* {\n")
		builder.WriteString(GetType(typ.Elem()))
		builder.WriteString(fmt.Sprintf("Methods(%d):\n", typ.NumMethod()))
		for i := 0; i < typ.NumMethod(); i++ {
			method := typ.Method(i)
			builder.WriteString(fmt.Sprintf("\tfunc %s%s\n", method.Name, "()"))
		}
		builder.WriteString("}\n")
		return builder.String()
	default:
		return "unkonw"
	}
}

func main() {
	var es = []interface{}{
		1, 1.1, "test", false,
		[2]int{1, 2}, []int{1, 2, 3}, map[int]string{1: "kk"},
		User{}, &User{},
	}

	// 打印每种变量的数据类型
	for _, e := range es {
		fmt.Println(GetType(reflect.TypeOf(e)))
	}
}
