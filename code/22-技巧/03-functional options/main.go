package main

import (
	"fmt"
	"time"
)

// Server 要有侦听的 IP 地址 Addr 和端口号 Port ，这两个配置选项是必填的
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConn  int
}

// NewDefaultServer 第一种需要有多种不同的创建不同配置 Server 的函数签名
func NewDefaultServer(addr string, port int) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
	return &Server{addr, port, "tcp", timeout, 100}, nil
}

// ServerBuilder 2.使用一个builder类来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	//其它代码设置其它成员的默认值
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxConn int) *ServerBuilder {
	sb.Server.MaxConn = maxConn
	return sb
}

func (sb *ServerBuilder) WithTimeOut(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) Build() Server {
	return sb.Server
}

func main() {
	sb := ServerBuilder{}
	server := sb.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(1024).
		WithTimeOut(30 * time.Second).
		Build()
	fmt.Println(server)

	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConn(1000))
	fmt.Println(s1)
	fmt.Println(s2)
}

// Option 4.Functional Options
type Option func(server *Server)

func ProtoPool(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}

func Timeout(time time.Duration) Option {
	return func(server *Server) {
		server.Timeout = time
	}
}

func MaxConn(maxConn int) Option {
	return func(server *Server) {
		server.MaxConn = maxConn
	}
}
func NewServer(add string, port int, options ...Option) (*Server, error) {
	srv := &Server{
		Addr:     add,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConn:  1000,
	}

	for _, option := range options {
		option(srv)
	}

	return srv, nil
}
