package main

import (
	"fmt"
	"os"
)

// os.O_WRONLY	只写
// os.O_CREATE	创建文件
// os.O_RDONLY	只读
// os.O_RDWR	读写
// os.O_TRUNC	清空
// os.O_APPEND	追加
func main() {
	file, err := os.OpenFile("./write.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open err", err)
		return
	}
	defer file.Close()

	// Write / WriteString
	str := "hello 迪迦"
	file.Write([]byte(str + "\n"))
	file.WriteString("hello 大古\n")
}
