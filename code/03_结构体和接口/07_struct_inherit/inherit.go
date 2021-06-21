package main

import "fmt"

// 模拟继承
// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type dog struct {
	Feet    int8
	*animal //通过嵌套匿名结构体实现继承
}

func (d *dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &dog{
		Feet: 4,
		animal: &animal{ //注意嵌套的是结构体指针
			name: "旺财",
		},
	}
	d1.wang() //旺财会汪汪汪~
	d1.move() //旺财会动！
}
