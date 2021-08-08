package main

import "fmt"

func main() {
	names := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println("请输入要查找的名字")
	var s string
	fmt.Scanln(&s)
	// 一
	for i := 0; i < len(names); i++ {
		if s == names[i] {
			fmt.Printf("找到%v 下标为%d", s, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Println("没有找到")
		}
	}

	//二
	index := -1
	for i := 0; i < len(names); i++ {
		if s == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v 下标为%d", s, index)
	} else {
		fmt.Println("没有找到")
	}
}
