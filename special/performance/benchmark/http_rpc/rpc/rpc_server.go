package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
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

		log.Println("new request...")
		rpc.ServeConn(conn) // 源码注释上说，应该异步调用的
	}
}
