package example2

import (
	"go-learn/core/net/grpc/example2/proto/hello"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestServerRun(t *testing.T) {
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, NewHelloService())

	l, err := net.Listen("tcp", ":" + ServerPortStr())
	if err != nil {
		t.Fatal(err)
	}
	if err = server.Serve(l); err != nil {
		t.Fatal(err)
	}
}
