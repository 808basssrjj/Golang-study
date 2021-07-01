package main

import (
	"fmt"
	"strconv"
)

// strconv包实现了基本数据类型和其字符串表示的相互转换
// Atoi()、Itia()、parse系列、format系列、append系列。
func main() {
	// 1.string与int类型转换
	// 字符串转int Atoi
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
	// int转字符串 Itoa
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	// 2.Parse系列函数 用于转换字符串为给定类型的值
	// 2.1ParseBool
	// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)
	// 2.2ParseInt
	i, err := strconv.ParseInt("-2", 10, 64)
	fmt.Println(i, err)
	// 2.3ParseUnit()
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(u, err)
	// 2.4ParseFloat()
	f, err := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f, err)

	// 3.Format系列函数   实现了将给定类型数据格式化为string类型数据的功能
	st1 := strconv.FormatBool(true)
	st2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	st3 := strconv.FormatInt(-2, 16)
	st4 := strconv.FormatUint(2, 16)
	fmt.Println(st1, st2, st3, st4)

}
