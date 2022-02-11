package _step

import (
	"go-learn/core/net/rpc/2step/proto/hello"
	"testing"
)

func TestClientCall(t *testing.T) {
	client, err := NewHelloClient("tcp", ":"+ServerPortStr())
	if err != nil {
		t.Fatal(err)
	}

	var (
		args  = &hello.Args{Value: "I'm client"}
		reply = &hello.Reply{}
	)
	if err = client.Hello(args, reply); err != nil {
		t.Fatal(err)
	}
	t.Log(reply)
}
