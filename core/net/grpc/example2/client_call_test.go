package example2

import (
	"context"
	"go-learn/core/net/grpc/example2/proto/hello"
	"testing"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient(ServerAddr())
	if err != nil {
		t.Fatal(err)
	}

	c, err := client.Hello(context.Background(), &hello.HelloArgs{ Value: "Hello! I'm client!" })
	if err != nil {
		t.Fatal(err)
	}

	reply, err := c.Recv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)
	reply, err = c.Recv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)

	if err = c.CloseSend(); err != nil {
		t.Fatal(err)
	}
}
