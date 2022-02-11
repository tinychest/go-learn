package _step

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestServerRun(t *testing.T) {
	err := rpc.RegisterName(ServiceName(), NewHelloService())
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":" + ServerPortStr())
	if err != nil {
		t.Fatalf("ListenTCP error: %s", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatalf("accept error: %s", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
