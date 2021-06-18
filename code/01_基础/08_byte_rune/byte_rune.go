package main

import "fmt"

func main() {
	//组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来
	s := "hello沙河"
	n := len(s)
	fmt.Println(n)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	// 修改字符串
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	s1 := "big"
	byteS1 := []byte(s1) // 强制转换为byte
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2) // 强制转换为rune切片
	runeS2[0] = '红'
	fmt.Println(string(runeS2))

	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("%T\n", f)
}
