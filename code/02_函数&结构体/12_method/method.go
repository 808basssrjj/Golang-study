package main

import "fmt"

type dog struct {
	name string
}
type person struct {
	name string
	age  int
}

func newdog(name string) *dog {
	return &dog{
		name: name,
	}
}
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// Go语言中的方法（Method）是一种作用于特定类型变量的函数。
// 这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
// 接收者表示的是调用该方法的具体类型变量,多用类型名小写首字母表示
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//     函数体
// }
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

// 2.1使用值接收者
// 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
// 在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
func (p person) old() {
	p.age++
}

// 2.2使用指针接收者
// 指针类型的接收者由一个结构体的指针组成，
// 由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
func (p *person) old2() {
	p.age++
}

// 什么时候用指针接收者?
// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

func main() {
	d1 := newdog("旺财")
	d1.wang()

	p1 := newPerson("张三", 17)
	fmt.Println(p1.age) //17
	p1.old()
	fmt.Println(p1.age) //18
	p1.old2()
	fmt.Println(p1.age) //18
}
