package helloworld

import (
	"net"
	"net/rpc"
	"testing"
)

func TestServerRun(t *testing.T) {
	err := rpc.RegisterName(ServiceName(), NewHelloService())
	if err != nil {
		t.Fatalf("register service faild: %s", err)
	}

	listener, err := net.Listen("tcp", ":" + ServerPortStr())
	if err != nil {
		t.Fatalf("ListenTCP error: %s", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		t.Fatalf("accept error: %s", err)
	}
	rpc.ServeConn(conn)
}
