package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("D:/b")
	if err == nil {
		fmt.Println("目录或文件存在")
	} else {
		fmt.Println(err)
	}
	if os.IsExist(err) {
		fmt.Println("目录或文件不存在")
	}
}
