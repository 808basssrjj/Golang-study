package main

import "fmt"

func main() {
	// 切片不存值，本质是一个框，框住了一块连续的内存
	// 切片属于引用类型，真正的数据保存在底层数组中

	//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合
	// var name []T			name:表示变量名T:表示切片中的元素类型
	// var s []int

	// 1.切片初始化
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)              //[]
	fmt.Println(b)              //[]
	fmt.Println(c)              //[false true]
	fmt.Println(a == nil)       //true
	fmt.Println(b == nil)       //false
	fmt.Println(c == nil)       //false

	// 2.长度和容量
	// 容量是底层数组从切片的第一个元素到最后以一个元素的数量
	fmt.Println(len(c))
	fmt.Println(cap(c))

	// 3.由数组得到切片 (左闭右开)
	arr := [...]int{1, 3, 5, 7, 9}
	s1 := arr[0:4]
	s2 := arr[:3] // 等同于 a[0:3]
	s3 := arr[:]  // 等同于 a[0:len(a)]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 4.make()函数创建切片
	// make([]T, size, cap)
	// T:切片的元素类型
	// size:切片中元素的数量
	// cap:切片的容量
	s := make([]int, 2, 10)
	fmt.Println(s)      //[0,0]
	fmt.Println(len(s)) //2
	fmt.Println(cap(s)) //10

	// 5.要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
	// 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil
	var a1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	a2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	a3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Println(a1, a2, a3)

	// 6.append()添加元素
	var sl1 []int
	sl1 = append(sl1, 1)       //[1]
	sl1 = append(sl1, 2, 3, 4) //[1,2,3,4]
	sl2 := []int{5, 6, 7}
	sl1 = append(sl1, sl2...) //[1,2,3,4,5,6,7]  ...拆分切片
	//每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	//当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“
	//扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值。
	//扩容策略 :
	//1.如果申请的容量大于原来两倍，那就直接扩容到新申请的容量
	//2.如果小于1024，直接两倍
	//3.如果打印1024,1.25倍
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%d:%v  len:%d  cap:%d  ptr:%p\n", i, numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	// 7.copy()复制切片 //由于切片是引用类型，修改一个会修改其他，所以要用copy
	//copy(destSlice, srcSlice []T)
	//srcSlice: 数据来源切片
	//destSlice: 目标切片
	srcSlice := []int{1, 2, 3}
	destSlice := make([]int, 3, 3)
	// destSlice = append(destSlice, 1)
	copy(destSlice, srcSlice)
	fmt.Println(srcSlice, destSlice)
	srcSlice[0] = 1000
	fmt.Println(srcSlice, destSlice)

	// 8.删除切片元素
	// 要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	delete := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	delete = append(delete[:2], delete[3:]...)
	fmt.Println(delete)

	// 练习
	var aa = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, i)
	}
	fmt.Println(aa)
}
