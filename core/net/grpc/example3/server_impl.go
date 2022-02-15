package example3

import (
	"fmt"
	"go-learn/core/net/grpc/example3/proto/hello"
)

type HelloService struct{}

func NewHelloService() hello.HelloServer {
	return new(HelloService)
}

func (hc *HelloService) Hello(stream hello.Hello_HelloServer) error {
	fmt.Println("a new call...")

	// 接收
	reply, err := stream.Recv()
	if err != nil {
		return err
	}
	fmt.Println("recv:", reply.Value)

	// 发送
	args := &hello.HelloReply{Value: "pong"}
	err = stream.Send(args)
	if err != nil {
		return err
	}
	fmt.Println("send:", args.Value)

	// 没有关闭方法
	return nil
}
