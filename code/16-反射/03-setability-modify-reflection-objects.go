package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 获取值和修改值

	fVar := 3.14
	v := reflect.ValueOf(fVar)
	fmt.Printf("canSet : %v canAddr : %v\n", v.CanSet(), v.CanAddr())

	vp := reflect.ValueOf(&fVar)
	// Elem() 可通过指针获取值
	fmt.Printf("canSet : %v canAddr : %v\n", vp.Elem().CanSet(), vp.Elem().CanAddr())
	// Set系列 可以修改值
	vp.Elem().SetFloat(2.333)
	fmt.Println(fVar)
}
