package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	// 注册路由
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(fmt.Errorf("ListenTCP error: %w", err))
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(fmt.Errorf("accept error: %w", err))
		}

		fmt.Println("new request...")
		rpc.ServeConn(conn) // 源码注释上说，应该异步调用的
	}
}
