package example2

import (
	"fmt"
	"go-learn/core/net/grpc/example2/proto/hello"
)

type HelloService struct{}

func NewHelloService() hello.HelloServer {
	return new(HelloService)
}

func (hc *HelloService) Hello(args *hello.HelloArgs, server hello.Hello_HelloServer) error {
	fmt.Println("a new call:", args.Value)

	err := server.Send(&hello.HelloReply{Value: "hello! I'm server!"})
	if err != nil {
		return err
	}
	err = server.Send(&hello.HelloReply{Value: "hello! I'm server! again!"})
	if err != nil {
		return err
	}
	return nil
}
