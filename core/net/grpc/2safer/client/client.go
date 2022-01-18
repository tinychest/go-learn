package main

import (
	"fmt"
	_safer2 "go-learn/core/net/grpc/2safer"
	"net/rpc"
)

var _ _safer2.HelloService = (*HelloClient)(nil)

type HelloClient struct {
	*rpc.Client
}

func (p *HelloClient) Hello(request string, reply *string) error {
	return p.Client.Call(_safer2.HelloServiceName+".Hello", request, reply)
}

func NewHelloClient(network, address string) (*HelloClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloClient{
		Client: c,
	}, nil
}

// 其中唯一的变化是 client.Call 的第一个参数用 HelloServiceName+".Hello" 代替了 "HelloService.Hello"
// 然而通过 client.Call 函数调用RPC方法依然比较繁琐，同时参数的类型依然无法得到编译器提供的安全保障
func main0() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(fmt.Errorf("dialing: %w", err))
	}

	var reply string
	err = client.Call(_safer2.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}
}

// 现在再也不会发生方法名或参数类型不匹配等低级的错误了
func main() {
	client, err := NewHelloClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var reply string
	if err = client.Hello("hello", &reply); err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
