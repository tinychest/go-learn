package _step

import "go-learn/core/net/rpc/2step/proto/hello"

// IHello 客户端 和 服务端 约定一致的接口，服务端需要实现该接口返回数据，客户端需要实现该接口获取数据
type IHello interface {
	Hello(request *hello.Args, reply *hello.Reply) error
}
