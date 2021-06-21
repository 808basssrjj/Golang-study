package main

import "fmt"

// 在Go语言中接口（interface）是一种类型，一种抽象的类型。
// 一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。
// 使用场景: 我不关心变量是什么类型,我只关心能调用它什么方法

// 定义
// type 接口类型名 interface{
//     方法名1( 参数列表1 ) 返回值列表1
//     方法名2( 参数列表2 ) 返回值列表2
//     …
// }

// 定义一个能叫的类型
type speaker interface {
	speak() //只要实现了speak方法的变量都是sperker类型
}

type dog struct{}
type cat struct{}
type person struct{}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}
func (c cat) speak() {
	fmt.Println("喵喵喵")
}
func (p person) speak() {
	fmt.Println("啊啊啊")
}

// 传谁打谁
// 需要一个特殊的类型
func da(x speaker) {
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)

	// 定义一个sperker类型
	var ss speaker
	ss = c1
	ss.speak()
	ss = d1
	ss.speak()
	ss = p1
	ss.speak()
}
