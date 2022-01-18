package main

import (
	_safer2 "go-learn/core/net/grpc/2safer"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func RegisterHelloService(svc _safer2.HelloService) error {
	return rpc.RegisterName(_safer2.HelloServiceName, svc)
}

func main() {
	err := RegisterHelloService(new(HelloService))
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
