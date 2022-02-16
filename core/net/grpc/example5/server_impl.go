package example5

import (
	"context"
	"fmt"
	"go-learn/core/net/grpc/example5/proto/hello"
)

type HelloServer struct{}

func NewHelloServer() hello.HelloServer {
	return new(HelloServer)
}

func (s *HelloServer) Hello(ctx context.Context, args *hello.HelloArgs) (*hello.HelloReply, error) {
	fmt.Println("a new call:", args.Value)
	return &hello.HelloReply{Value: "Hello! I'm server!"}, nil
}

