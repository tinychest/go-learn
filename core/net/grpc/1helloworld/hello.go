package helloworld

import (
	"go-learn/core/net/grpc/proto/hello"
)

// HelloService is rpc server obj
type HelloService struct{}

// Hello 方法的输入参数和输出的参数均改用 Protobuf 定义的 String 类型表示。
// 因为新的输入参数为结构体类型，因此改用指针类型作为输入参数，函数的内部代码同时也做了相应的调整。
func (p *HelloService) Hello(request *hello.String, reply *hello.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
