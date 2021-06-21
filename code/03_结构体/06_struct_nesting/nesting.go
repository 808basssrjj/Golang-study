package main

import "fmt"

// 嵌套结构体
type address struct {
	province string
	city     string
}

type workplace struct {
	province string
	city     string
}

type persion struct {
	name    string
	age     int
	address //匿名字段
	workplace
}

type company struct {
	name    string
	address address
}

func main() {
	// 1.基本使用
	p1 := persion{
		name: "张三",
		age:  18,
		address: address{
			province: "浙江",
			city:     "诸暨",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)

	// 2.嵌套匿名字段
	// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
	// fmt.Println(p1.city)

	// 3.嵌套结构体的字段名冲突
	// 为了避免歧义需要通过指定具体的内嵌结构体字段名
	fmt.Println(p1.address.city)
	fmt.Println(p1.workplace.city)

}
