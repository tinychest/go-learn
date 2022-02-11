package _step

import (
	"fmt"
	"go-learn/core/net/rpc/2step/proto/hello"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type helloClient struct {
	*rpc.Client
}

func NewHelloClient(network, address string) (IHello, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		panic(fmt.Errorf("net.Dial: %w", err))
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	return &helloClient{
		Client: client,
	}, nil
}

func (p *helloClient) Hello(args *hello.Args, reply *hello.Reply) error {
	return p.Client.Call(ServiceName()+"."+ServiceFuncName(), args, reply)
}
