package main

import "fmt"

// 1.使用空接口实现可以接收任意类型的函数参数。
func show(x interface{}) {
	fmt.Printf("type:%T value:%v\n", x, x)
}
func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	// 1.
	aa := "迪迦"
	bb := 2000
	show(aa)
	show(bb)

	// 2.空接口作为map的值
	var interfaceMap = make(map[string]interface{})
	interfaceMap["name"] = "张三"
	interfaceMap["age"] = 18
	interfaceMap["isdie"] = false
	fmt.Println(interfaceMap)

}
