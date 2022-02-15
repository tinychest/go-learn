package example2

import (
	"go-learn/core/net/grpc/example2/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HelloClient struct{}

func NewHelloClient(target string) (hello.HelloClient, error) {
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return hello.NewHelloClient(conn), nil
}
