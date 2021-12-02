package main

import (
	"go-learn/core/net/rpc/1helloworld"
	"log"
	"net"
	"net/rpc"
)

/* rpc 服务端 */

func main() {
	_ = rpc.RegisterName("HelloService", new(helloworld.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error", err)
	}

	rpc.ServeConn(conn)
}
