package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) PrintName() {
	fmt.Println(u.Name)
}

func (u User) PrintAge() {
	fmt.Println(u.Age)
}

type Aggregator func(int, int) int

var (
	user = User{
		Name: "迪迦",
		Age:  18,
	}
	add Aggregator = func(a int, b int) int {
		return a + b
	}
	sub Aggregator = func(a int, b int) int {
		return a - b
	}
)

// 通过反射获得variable的各种信息
func inspect(variable interface{}) {
	v := reflect.ValueOf(variable)
	t := reflect.TypeOf(variable)

	if t.Kind() == reflect.Struct {
		fmt.Printf("--------- field: %d ---------\n", t.NumField())
		// 1.通过for循环遍历结构体的所有字段信息
		for idx := 0; idx < t.NumField(); idx++ {
			fieldType := t.Field(idx)
			fieldValue := v.Field(idx)
			fmt.Printf("\t %v = %v\n", fieldType.Name, fieldValue)
			fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
				fieldType.Name, fieldType.Index, fieldType.Type, fieldType.Tag.Get("json"))
		}
		// 2.通过字段名获取指定结构体字段信息
		//if ageField, ok := t.FieldByName("Age"); ok {
		//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
		//		ageField.Name, ageField.Index, ageField.Type, ageField.Tag.Get("json"))
		//}

		fmt.Printf("--------- method: %d ---------\n", t.NumMethod())
		for idx := 0; idx < t.NumMethod(); idx++ {
			methodType := t.Method(idx)
			fmt.Printf("\t method_name=%s input_num=:%d output_num:%d\n",
				methodType.Name, methodType.Type.NumIn(), methodType.Type.NumOut())
		}

	} else if t.Kind() == reflect.Func {
		fmt.Printf("this function has %d inputs and %d outputs\n", t.NumIn(), t.NumOut())
	}
}

func main() {
	inspect(user)
	inspect(add)
	inspect(sub)
}
