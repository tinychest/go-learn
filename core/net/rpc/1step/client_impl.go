package helloworld

import (
	"fmt"
	"net/rpc"
)

// 客户端封装具体的远程调用方式是具有重大意义的，否则，调用 Call 将非常生硬，还容易因为敲错单词导致报错

type helloClient struct{}

// 要求 helloClient 实现 IHello 接口

func NewHelloClient() IHello {
	return new(helloClient)
}

func (c *helloClient) Hello(args *Args, reply *Reply) error {
	client, err := rpc.Dial("tcp", ServerAddr())
	if err != nil {
		return fmt.Errorf("dialing err: %w", err)
	}
	return client.Call(ServiceName()+"."+ServiceFuncName(), args, reply)
}
