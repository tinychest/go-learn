package example1

import (
	"context"
	"go-learn/core/net/grpc/example1/proto/hello"
)

type HelloService struct {}

func NewHelloService() hello.HelloServer {
	return new(HelloService)
}

func (s HelloService) Hello(ctx context.Context, args *hello.HelloArgs) (*hello.HelloReply, error) {
	return &hello.HelloReply{Value: args.Value}, nil
}
