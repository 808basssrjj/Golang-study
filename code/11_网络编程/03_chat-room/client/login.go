package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_code/03_chat/common"
	"net"
)

// login 登录验证
func login(id int, pwd string) (err error) {
	//fmt.Printf("id:%d password:%s\n",id,pwd)
	//return nil

	// 1.连接客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()

	// 2.通过conn向客户端发送信息
	var msg message.Message
	loginMsg := message.LoginMsg{UserId: id, UserPwd: pwd}
	data, err := json.Marshal(loginMsg) //序列化
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	msg.Data = string(data)         //序列化后赋给meg
	msg.Type = message.LoginMegType //序列化后赋给meg

	// 序列化msg
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}

	// 发送data
	// 先发送长度 要切片
	pgkLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pgkLen)
	n, err := conn.Write(buf[:4]) //发送长度
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) failed", err)
		return
	}
	// fmt.Printf("客户端发送数据的长度%v\n 内容是%v\n",len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data) //发送长度
	if err != nil {
		fmt.Println("conn.Write(data) failed", err)
		return
	}

	// 处理服务器返回的消息
	msg, err = readPkg(conn)
	
	if err != nil {
		fmt.Println()
	}

	var loginResMsg message.LoginRes
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		fmt.Println("登录成功！")
	} else if loginResMsg.Code == 500 {
		fmt.Println(loginResMsg.Error)
	}
	return
}
