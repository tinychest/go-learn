package helloworld

import (
	"testing"
)

func TestClientCall(t *testing.T) {
	var (
		args  = &Args{Value: "I'm client"}
		reply = &Reply{}
	)
	if err := NewHelloClient().Hello(args, reply); err != nil {
		t.Fatal(err)
	}
	t.Log("success rpc result:", reply)
}
