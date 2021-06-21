package main

import "fmt"

//类型断言

// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
// x.(T)
// 其中：
// - x：表示类型为`interface{}`的变量
// - T：表示断言`x`可能是的类型。
// 该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

//
func assert(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("x is a string value:", v)
	case int:
		fmt.Println("x is a	int value:", v)
	case bool:
		fmt.Println("x is a bool value:", v)
	default:
		fmt.Println("unsupport type！")
	}
}
func main() {
	// 断言1
	var x interface{}
	s := "hello 迪迦"
	x = s
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("断言失败")
	}

	// 断言2
	assert(100)
	assert(true)
	assert("ssss")

}
