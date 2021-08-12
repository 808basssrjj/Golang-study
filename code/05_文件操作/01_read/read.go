package main

import (
	"fmt"
	"io"
	"os"
)

// 文件在程序中是以流的形式来操作的
// 流:数据在数据源(文件)和程序(内存)之间经历的路径
// 输入流:读文件
// 输出流:写文件


// read读取}
func main() {
	// 打开文件
	fileObj, err := os.Open("./01_file.go")
	if err != nil {
		fmt.Println("open file err", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()

	// 读文件Read
	// 定义 func (f *File) Read(b []byte) (n int, err error)
	// 指定读取长度
	// var tmp = make([]byte, 128)
	// n, err := fileObj.Read(tmp)
	var tmp [128]byte
	n, err := fileObj.Read(tmp[:])
	if err == io.EOF {
		println("读取完成")
		return
	}
	if err != nil {
		fmt.Println("read file err", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))
}
