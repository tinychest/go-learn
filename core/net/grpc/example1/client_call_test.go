package example1

import (
	"context"
	"go-learn/core/net/grpc/config"
	"go-learn/core/net/grpc/example1/proto/hello"
	"testing"
	"time"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient(config.ServerAddr())
	if err != nil {
		t.Fatal(err)
	}

	// 对 rpc 调用进行超时设定
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var args = &hello.HelloArgs{Value:"I'm client"}

	reply, err := client.Hello(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)
}
