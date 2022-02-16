package example6

import (
	"context"
	"go-learn/core/net/grpc/config"
	"go-learn/core/net/grpc/example3/proto/hello"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient(config.ServerAddr())
	if err != nil {
		t.Fatal(err)
	}
	// 连接 和 发送 head
	headMD := metadata.Pairs("key", "value")
	ctx := metadata.NewOutgoingContext(context.Background(), headMD)
	stream, err := client.Hello(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("conn success...")
	t.Log("send head:", headMD["key"])

	// 发送
	args := &hello.HelloArgs{Value: "ping"}
	if err = stream.Send(args); err != nil {
		t.Fatal(err)
	}
	t.Log("send:", args.Value)

	// 接收 head（Client stream function）
	header, err := stream.Header()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("recv head:", header.Get("key"))
	// 接收
	reply, err := stream.Recv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("recv:", reply.Value)

	// 结束
	if err = stream.CloseSend(); err != nil {
		t.Fatal(err)
	}
}
