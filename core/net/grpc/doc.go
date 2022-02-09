package grpc

/*
【gRPC】
[参考]：https://chai2010.cn/advanced-go-programming-book/ch4-rpc/readme.html
[参考]：https://grpc.io/docs/what-is-grpc/introduction/
[参考]: https://grpc.io/docs/what-is-grpc/core-concepts/

[参考]: https://mp.weixin.qq.com/s/5vbhRdqGAiQDFo_9YeiS2g


[概念]
Like many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types.
By default, gRPC uses protocol buffers as the Interface Definition Language (IDL) for describing both the service interface and the structure of the payload messages.
It is possible to use other alternatives if desired.

不说的那么细节、专业，gRPC 是 Google 开发的一个高性能开源的 RPC 框架，基于 HTTP/2，同时支持大多数流行的编程语言。

gRPC 不像传统的 RPC 方法调用，不仅支持 一请求一响应；还拓展了流式的 一请求多响应 和 多请求多响应

[IDL（Interface Description Language）]
接口描述语言，例如 proto buffer。一些跨平台的 RPC 框架可以根据 IDL，在编译时期使用代码生成器生成 stub 代码。

[HTTP/2]
> 特别提及的基础协议 HTTP/2
HTTP/2 采用二进制格式传输协议，HTTP/1.x 是文本格式。
HTTP/2 支持通过一个连接发送多个并发请求。
HTTP/2 中，服务器可以对客户端的一个请求给予多个响应（服务端推送）；这是 HTTP/1.x 做不到的。
HTTP/2 对消息头进行了压缩，能够节省网络流量。

[Protocol Buffer]
详见，同级目录的 proto 包

【实战】
1.定义 proto 数据结构 → 使用 gRPC 提供的 protocol buffer 编译插件，生成相应的 Go 代码
2.定义 gRPC 接口
3.开发接口实现（服务端）
4.开发接口调用（客户端）
5.引入客户端代码实现调用服务端指定方法（RPC）

[其他]
https://pkg.go.dev/encoding/gob
*/
