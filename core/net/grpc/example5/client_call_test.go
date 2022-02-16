package example5

import (
	"context"
	"go-learn/core/net/grpc/config"
	"go-learn/core/net/grpc/example5/proto/hello"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestClientRun(t *testing.T) {
	client, err := NewHelloClient(config.ServerAddr())
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	headMD := metadata.Pairs("token", "administrator")
	ctx = metadata.NewOutgoingContext(context.Background(), headMD)

	args := &hello.HelloArgs{Value: "Hello! I'm client!"}

	reply, err := client.Hello(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply.Value)
}
