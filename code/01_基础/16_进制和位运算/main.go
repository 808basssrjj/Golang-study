package main

import "fmt"

func main() {
	// 八进制0 开头
	//i1 := 021
	//fmt.Println(i1)
	// 十六进制 0x开头
	//i2:=0x11
	//fmt.Println(i2)

	// 1.其他进制转十进制
	//i2 := 02456 //6+5*8+4*64+2*(8**3)=1326
	//i3 := 0xA45 //5+4*16+10*16*16 = 2629
	//fmt.Println(i2)
	//fmt.Println(i3)

	// 2.十进制转其他
	// 将该数不断除以n,知道商为O,然后将每步的余数倒过来

	// 3.二进制(binary)转其他
	// 二 -> 八    每三位为一组,转换为对应10进制,然后拼接 11010101 = 0 325
	// 二 -> 十六  每四位为一组,转换为对应10进制,然后拼接 11010101 = 0x D5

	// 4.其他转二进制
	// 八  -> 二   每一位转成三位的二进制,拼接
	// 十六 -> 二   每一位转成四位的二进制,拼接

	// 5.原码,反码,补码
	// 5.1 有符号:二进制的最高位代表符号为 0:正数 1:负数
	// 5.2 正数的原码,反码,补码都一样
	// 5.3 负数的反码=原码符号位不变,其他位取反
	// 5.4 负数的补码=它的反码+1
	// 5.5 0的反码,补码都是0
	// 5.6 计算机运算时,都是补码来运算的

	// 6.位运算
	// 6.1 & 按位与   两位全为1,结果为1,否则0
	// 6.2 | 按位或   两位有一个,结果为1,否则0
	// 6.3 ^ 按位异或  两位一个1一个0,结果为1,否则0
	// 要补码运算
	fmt.Println(2 & 3) //0000 0010 & 0000 0011 = 0000 0010 = 2
	fmt.Println(2 | 3) //0000 0010 | 0000 0011 = 0000 0011 = 3
	fmt.Println(2 ^ 3) //0000 0010 ^ 0000 0011 = 0000 0001 = 1
	//   -2反码1111 1101 -2补码1111 1110
	fmt.Println(-2 & 3) //1111 1110 & 0000 0011 = 0000 0010 = 2
	fmt.Println(-2 | 3) //1111 1110 | 0000 0011 = 1111 1111  (负数的补码 要返回成原码) (反))1111 1110 => (原)1000 0001 = -1
	fmt.Println(-2 ^ 3) //1111 1110 ^ 0000 0011 = 1111 1101  (负数的补码 要返回成原码) (反))1111 1100 => (原)1000 0011 = -3
	fmt.Println("---------------")

	// 6.4 右移>> 低位溢出,符号位不变,并用符号位补位溢出的高位
	// 6.5 左移<< 符号位不变,低位补0
	var a int = 1 >> 2  //0000 0001 => 0000 0000 = 0
	var b int = -1 >> 2 //1111 1111 => 1111 1111   (负数的补码 要返回成原码) => (反)1111 1110 => (原)1000 0001 = -1
	var c int = 1 << 2  //0000 0001 => 0000 0100 = 4
	var d int = -1 << 2 //1111 1111 => 1111 1100   (负数的补码 要返回成原码) => (反)1111 1011 => (原)1000 0100 = -4
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
