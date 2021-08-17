package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_code/03_chat/common"
	"net"
)

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
