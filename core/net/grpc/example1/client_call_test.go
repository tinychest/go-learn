package example1

import (
	"context"
	"go-learn/core/net/grpc/example1/proto/hello"
	"testing"
	"time"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient(ServerAddr())
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var args = &hello.HelloArgs{Value:"I'm client"}

	reply, err := client.Hello(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)
}
