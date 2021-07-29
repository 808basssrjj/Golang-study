package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func (s *Student) DoHomework(number int) {
	fmt.Printf("%s is doing homework, %d\n", s.name, number)
}
func main() {
	// 访问结构体方法
	// 可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息
	s := Student{name: "李四"}
	v := reflect.ValueOf(&s)
	methodV := v.MethodByName("DoHomework")
	if methodV.IsValid() {
		in := []reflect.Value{reflect.ValueOf(55)}
		methodV.Call(in)
	}
}
