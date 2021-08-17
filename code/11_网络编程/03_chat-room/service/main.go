package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_code/03_chat/common"
	"io"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 循环读取客户端发送的信息
	//buf := make([]byte, 8096)
	for {
		msg, err := readPkg(conn)
		//n, err := conn.Read(buf[:4])
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了,服务端也退出")
				return
			} else {
				fmt.Println("conn.Write(bytes) failed", err)
				return
			}
		}
		//fmt.Printf("读到的数据长度:%d\n", n)
		// fmt.Println("msg=", msg)
		err = serverProcessMsg(conn, &msg)
		if err != nil {
			return
		}
	}
}

// 读取客户端数据
func readPkg(conn net.Conn) (msg message.Message, err error) {
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

// 发送客户端数据
func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送长度给对方
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4]) //发送长度
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) failed", err)
		return
	}

	// 发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("writer failed err=", err)
		return
	}
	return
}

// 根据客户端消息不同，决定调用哪个函数
func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMegType:
		// 处理登录
		err = serverProcessLogin(conn, msg)
	case message.RegisterMsgType:
		//处理注册
		// err = serverProcessRegister(conn, msg)
	default:
		fmt.Println("消息类型不存在！")
	}
	return
}

// 处理登录消息
func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
	// 1.取出msg中的Data ,并发序列化为LoginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("unmarshal failed err=", err)
		return
	}

	var resMsg message.Message
	var loginResMsg message.LoginRes
	// 判断合法性
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
		loginResMsg.Code = 200
	} else {
		loginResMsg.Code = 500
		loginResMsg.Error = "用户不存在"
	}

	//3.将loginResMsg序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("marshal failed err=", err)
		return
	}

	// 4.将序列化后的数据给resMsg.Data
	resMsg.Data = string(data)
	// 5.将resMsg序列化后发送给客户端
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal failed err=", err)
	}
	err = writePkg(conn, data) //发送
	return
}

// 处理注册消息
// func serverProcessRegister(conn net.Conn, msg *message.Message) (err error) {
// }

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
