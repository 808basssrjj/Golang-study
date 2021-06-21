package main

import "fmt"


type student struct {
	name string
	age  int
}

// 构造函数  约定成俗new开头
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func main() {
	s1 := newStudent("张三", 18)
	fmt.Printf("%#v\n", s1)
}
