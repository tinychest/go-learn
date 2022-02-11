package _step

import (
	"go-learn/core/net/rpc/2step/proto/hello"
	"net/rpc"
)

type helloClient struct {
	*rpc.Client
}

func NewHelloClient(network, address string) (IHello, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &helloClient{
		Client: c,
	}, nil
}

func (p *helloClient) Hello(args *hello.Args, reply *hello.Reply) error {
	return p.Client.Call(ServiceName()+"."+ServiceFuncName(), args, reply)
}
