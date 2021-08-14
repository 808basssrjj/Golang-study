package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1.将interface{}转为reflect.Value   rVal:=reflect.ValueOf(B)
	// 2.将relfect.Valueleix 转换为interface{}  iVal:=reflect.Interface{}
	// 3.interface{} 转为原来的变量类型 使用类型断言 v:=iVal.(Stu)

	floatVar := 3.14
	v := reflect.ValueOf(floatVar)
	fmt.Println(v.Float() + 1.2) //获取float值
	newFloat := v.Interface().(float64) // 转为原来的变量类型
	fmt.Println(newFloat + 1.2)

	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	v = reflect.Append(v, reflect.ValueOf(4))
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 5)
	fmt.Println(newSlice)
}
