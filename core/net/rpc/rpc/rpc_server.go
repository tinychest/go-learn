package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func main() {
	// 注册路由
	_ = rpc.RegisterName("hello", new(Hello))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error", err)
		}

		fmt.Println("new request...")
		rpc.ServeConn(conn) // 源码注释上说，应该异步调用的
	}
}
