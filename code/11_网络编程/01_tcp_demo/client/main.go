package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("dial failed", err)
		return
	}
	defer conn.Close() // 关闭连接

	// 2.利用连接进行数据发送和接受
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			// fmt.Println("", err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("rece failed", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
