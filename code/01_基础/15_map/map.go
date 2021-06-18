package main

import "fmt"

func main() {
	//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	//map[KeyType]ValueType cap  TypeKeyType:表示键的类型。 ValueType:表示键对应的值的类型。 cap：表示容量，可以不传
	//map类型的变量默认初始值为nil，需要使用make()函数来分配内存

	// 1.基本使用
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 80
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("%T\n", scoreMap)

	// 2.声明时填充元素
	userInfo := map[string]string{
		"username": "张三",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 3.判断键是否存在
	// 如果key存在ok为true,value为对应的值；不存在ok为false,value为值类型的零值
	value, ok := scoreMap["王五"]
	if !ok {
		fmt.Println("查无此人")
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}

	// 4。遍历
	for k, v := range userInfo {
		fmt.Println(k, v)
	}

	// 5.delete删除
	scoreMap["王五"] = 10
	delete(scoreMap, "王五")
	fmt.Println(scoreMap)

	// 6.元素为map类型的切片
	mapSlice := make([]map[int]string, 10, 10)
	// 对内部map初始化
	mapSlice[0] = make(map[int]string, 2)
	mapSlice[0][10] = "迪迦"
	mapSlice[0][20] = "戴拿"
	mapSlice[1] = map[int]string{}
	mapSlice[1][10] = "赛文"
	fmt.Println(mapSlice)

	// 7.值为切片类型的map
	sliceMap := make(map[string][]int, 10)
	sliceMap["name"] = []int{10, 20, 30}
	fmt.Println(sliceMap)

}
