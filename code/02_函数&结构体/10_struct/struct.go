package main

import "fmt"

// Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，
// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。
// Go语言中通过struct来实现面向对象。
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

// type 类型名 struct {
//     字段名 字段类型
//     字段名 字段类型
//     …
// }
type person struct {
	name  string
	age   int
	hobby []string
}

// go中函数传参永远是拷贝
func f1(x person) {
	x.age = 1
}
func f2(x *person) {
	// (*x).age = 1
	x.age = 1
}

type test struct {
	a int8
	b int8
	c int8
}

func main() {
	// 1.实例化
	// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	var p1 person
	p1.name = "张三"
	p1.age = 18
	p1.hobby = []string{"篮球", "足球"}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Printf("%T\n", p1)

	// 2.匿名结构体  多用于临时场景
	var user struct {
		name string
		age  int
	}
	user.name = "迪迦"
	user.age = 10000000000000000

	// 3.struct是值类型
	f1(p1)
	fmt.Println(p1)
	f2(&p1)
	fmt.Println(p1)

	// 4.1创建指针类型结构体
	// 通过使用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	p2.name = "李青"             // 这是语法糖写法 = (*p2).name = "李青"
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	fmt.Printf("%p\n", p2)
	// 4.2取结构体的地址实例化  (常用)
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	// 5.初始化
	// 键值对初始化
	var p4 = person{
		name: "亚索",
		age:  20,
	}
	// 值的列表初始化
	var p5 = &person{
		"提莫",
		20,
		[]string{"a", "b"},
	}
	fmt.Printf("p4=%#v\n", p4)
	fmt.Printf("p5=%#v\n", p5)

	// 6.结构体内存布局
	// 结构体占用一块连续的内存
	t := test{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(t.a))
	fmt.Printf("%p\n", &(t.b))
	fmt.Printf("%p\n", &(t.c))

	// 7.结构体的匿名字段
	// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
	// type Person struct {
	// 	string
	// 	int
	// }
	// 这里匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
}
