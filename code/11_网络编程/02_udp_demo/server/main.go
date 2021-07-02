package main

import (
	"fmt"
	"net"
)

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) //接收数据
		if err != nil {
			fmt.Println("read udp failed", err)
			return
		}
		fmt.Printf("接收到数据:%v addr:%v", string(data[:n]), addr)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed err:%v\n", addr, err)
			return
		}
	}
}
