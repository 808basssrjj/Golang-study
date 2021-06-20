package main

import (
	"fmt"
	"io"
	"os"
)

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
