package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入  可输入空格
func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")
	s, _ = reader.ReadString('\n')
	fmt.Println("输入的是：", s)
}

// bufio读取
func main() {
	打开文件
	fileObj, err := os.Open("./01_file.go")
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()

	// 创建读文件的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			println("读取完成")
			return
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Print(line)
	}

	// useBufio()
}
