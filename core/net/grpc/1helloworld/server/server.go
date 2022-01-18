package main

import (
	"fmt"
	"go-learn/core/net/grpc/1helloworld"
	"net"
	"net/rpc"
)

/* rpc 服务端 */

func main() {
	_ = rpc.RegisterName("HelloService", new(helloworld.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(fmt.Errorf("ListenTCP error: %w", err))
	}

	conn, err := listener.Accept()
	if err != nil {
		panic(fmt.Errorf("accept error: %w", err))
	}

	rpc.ServeConn(conn)
}
