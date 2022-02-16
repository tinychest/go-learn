package example5

import (
	"context"
	"errors"
	"go-learn/core/net/grpc/config"
	"go-learn/core/net/grpc/example5/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"testing"
)

func Auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md["token"]) == 0 {
		return nil, errors.New("请求登录")
	}
	if md["token"][0] != "administrator" {
		return nil, errors.New("伪造登录信息")
	}

	return handler(ctx, req)
}

func TestServerRun(t *testing.T) {
	authServerOption := grpc.ChainUnaryInterceptor(Auth)

	server := grpc.NewServer(authServerOption)

	hello.RegisterHelloServer(server, NewHelloServer())

	l, err := net.Listen("tcp", ":"+config.ServerPortStr())
	if err != nil {
		t.Fatal(err)
	}

	if err = server.Serve(l); err != nil {
		t.Fatal(err)
	}
}
