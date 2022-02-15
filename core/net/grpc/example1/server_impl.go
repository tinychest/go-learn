package example1

import (
	"context"
	"fmt"
	"go-learn/core/net/grpc/example1/proto/hello"
)

type HelloService struct{}

func NewHelloService() hello.HelloServer {
	return new(HelloService)
}

func (s *HelloService) Hello(ctx context.Context, args *hello.HelloArgs) (*hello.HelloReply, error) {
	fmt.Println("a new call:", args.Value)
	return &hello.HelloReply{Value: "Hello! I'm server!"}, nil
}
