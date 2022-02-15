package example3

import (
	"context"
	"go-learn/core/net/grpc/example3/proto/hello"
	"testing"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient(ServerAddr())
	if err != nil {
		t.Fatal(err)
	}
	stream, err := client.Hello(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("conn success...")

	// 发送
	args := &hello.HelloArgs{Value: "ping"}
	if err = stream.Send(args); err != nil {
		t.Fatal(err)
	}
	t.Log("send:", args.Value)

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
