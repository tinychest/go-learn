package main

import (
	"fmt"
	over "go-learn/core/net/grpc/3over_lang"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	err := rpc.RegisterName(over.HelloServiceName, new(HelloService))
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(fmt.Errorf("accept error: %w", err))
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
