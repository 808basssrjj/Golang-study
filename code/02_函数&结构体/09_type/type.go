package main

import "fmt"

// 类型别名和自定义类型

// 1.自定义类型type后跟类型
//将MyInt定义为int类型
type Myint int

// 2.类型别名
// 类型别名只会在代码中存在，编译完成时并不会有
type yourInt = int
// type byte = uint8
// type rune = int32

func main() {
	var n Myint
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)	//main.Myint

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)	//int
}
