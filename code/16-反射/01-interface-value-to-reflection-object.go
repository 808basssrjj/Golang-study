package main

import (
	"fmt"
	"reflect"
)


func printMete (obj interface{}){
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//在反射中关于类型还划分为两种：类型（Type）和种类（Kind）
	//当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）
	kind :=t.Kind()
	name :=t.Name()
	fmt.Printf("type: %s value:%s kind: %s name:%s\n",t,v,kind,name)
}

type handler func(int,int)int

func main() {
	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name string
		pages int
	}
	sum := func(a,b int) int{
		return a+b
	}
	var sub handler =  func(a,b int)int {
		return a-b
	}

	printMete(intVar)
	printMete(stringVar)
	printMete(book{"go语言实战",100})
	printMete(sum)
	printMete(sub)
}