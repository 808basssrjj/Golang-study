package main

import "fmt"

func main() {
	// 需要指定长度和类型
	// var 数组变量名 [元素数量]T
	// var a [3]int

	// 1.数组初始化
	var testArray [3]int                          //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                   //使用指定的初始值完成初始化
	var cityArray = [...]string{"北京", "上海", "深圳"} //根据初始值的个数自行推断数组的长度
	fmt.Println(testArray)                        //[0 0 0]
	fmt.Println(numArray)                         //[1 2 0]
	fmt.Println(cityArray)                        //[北京 上海 深圳]
	a := [...]int{1: 1, 3: 5}                     //指定索引值的方式来初始化数组
	fmt.Println(a)

	// 2.遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	for index, value := range cityArray {
		fmt.Println(index, value)
	}

	// 3.二维数组
	b := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(b)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(b[2][1]) //支持索引取值:重庆
	for _, v1 := range b {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	a1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a1); i++ {
		for j := 0; j < i+1; j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}
}
