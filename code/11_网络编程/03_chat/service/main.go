package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_code/海量用户通讯系统/common"
	"io"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 循环读取客户端发送的信息
	//buf := make([]byte, 8096)
	for {
		msg, err := readPgk(conn)
		//n, err := conn.Read(buf[:4])
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出了,服务端也退出")
				return
			}else{
				fmt.Println("conn.Write(bytes) failed", err)
				return
			}
		}
		//fmt.Printf("读到的数据长度:%d\n", n)
		fmt.Println("msg=",msg)
	}
}

// 读取客户端数据
func readPgk(conn net.Conn) (msg message.Message, err error) {
	buf := make([]byte, 8096)
	//conn.Read 只有在没有关闭的情况下,才会阻塞
	n, err := conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pgk header failed")
		return
	}

	var pgkLen uint32
	pgkLen = binary.BigEndian.Uint32(buf[:4])
	// 根据pgkLen读取内容
	n, err = conn.Read(buf[:pgkLen])
	if n != int(pgkLen) || err != nil {
		//err = errors.New("read pgk body failed")
		return
	}

	// 反系列化成messqge.Messqge
	err = json.Unmarshal(buf[:pgkLen], &msg)
	if err != nil {
		fmt.Println("Unmarshal failed", err)
		return
	}

	return
}

func main() {
	// 1.启动服务
	fmt.Println("客户端在8889端口监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net listen err,", err)
	}

	// 2.建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,", err)
		}

		// 一旦链接成功,则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
