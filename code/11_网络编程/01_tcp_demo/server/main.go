package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端

// 连接处理
func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	// 针对当前连接做数据的发送和接受
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from conn failed", err)
			break
		}
		recv := string(buf[:n])
		fmt.Printf("接收到数据:%s\n", recv)
		conn.Write([]byte("ok")) // 发送数据
	}

}

func main() {
	// 1.启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}

	for {
		// 2.建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			continue
		}

		// 3.启动一个goroutine去处理连接
		go process(conn)
	}

}
