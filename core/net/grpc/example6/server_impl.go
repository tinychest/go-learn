package example6

import (
	"fmt"
	"go-learn/core/net/grpc/example3/proto/hello"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type HelloService struct{}

func NewHelloService() hello.HelloServer {
	return new(HelloService)
}

func (hc *HelloService) Hello(stream hello.Hello_HelloServer) error {
	fmt.Println("a new call...")

	// 接收 head
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "loss header")
	}
	fmt.Println("recv head:", md["key"])
	// 接收
	reply, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "recv error: %s", err)
	}
	fmt.Println("recv:", reply.Value)

	// 发送 head（Server stream function）
	headMD := metadata.Pairs("key", "value")
	// 该写法只在服务端生效
	// if err = grpc.SendHeader(stream.Context(), headMD); err != nil {
	// 	return status.Errorf(codes.Unknown, "send head error: %s", err)
	// }
	if err = stream.SendHeader(headMD); err != nil {
		return status.Errorf(codes.Unknown, "send head error: %s", err)
	}
	fmt.Println("send head:", headMD["key"])
	// 发送
	args := &hello.HelloReply{Value: "pong"}
	err = stream.Send(args)
	if err != nil {
		return status.Errorf(codes.Unknown, "send error: %s", err)
	}
	fmt.Println("send:", args.Value)

	// 没有关闭方法
	return nil
}
