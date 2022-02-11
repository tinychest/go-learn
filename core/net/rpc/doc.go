package rpc

/*
【RPC（Remote Procedure Call）】

[参考]:
https://zhuanlan.zhihu.com/p/187560185
https://mp.weixin.qq.com/s/5vbhRdqGAiQDFo_9YeiS2g

[简介]
直译远程过程调用，也叫远程方法调用。
它表示服务器 A 中的程序希望调用另外一台服务器 B 中的程序方法，因为不在一个内存空间中，所以只能通过网络来实现。

[细说]
RPC 采用 客户端/服务端 模式，它的三个要素是：通信协议（HTTP）、寻址（IP + PORT）、数据的序列化方式（JSON）

B 需要监听指定的网络端口，监听到请求后建立连接，然后，按照约定的网络传输协议，
接受参数，用于调用约定方法，然后将结果作为响应写回去；
对 A 来说，需要知道 B 的网络地址和监听的端口，才能主动发起请求、建立连接，进行数据交换。

RPC 框架是为了让 RPC 的实现更加简单、透明，它封装了上述细节。
这样调用者可以像调用本地接口一样调用远程的服务提供者，而不需要关注底层通信细节和调用过程。

[RPC 和 HTTP 有什么区别？]
这样的问法本身是不严谨的，HTTP 是通信协议，RPC 是远程调用方案，它的概念应该是这样的：
基于 HTTP 实现的 Restful JSON 数据格式的 接口调用，可以被称为是一种 RPC 实现。

虽然你说 HTTP 和 RPC 可能大家知道意思，但是还是应该专业一些，HTTP Restful 和 gRPC，这才是你大多数希望表达的正确意思。

说回来，就以这里上下文背景的，也是我们熟知的 HTTP Restful 形式的 RPC 和 gRPC 之间的一些区别：
传输效率：HTTP 1.1，通常包含很多包含很多无用的内容、直接基于 TCP 协议的 RPC，可以自定义上层协议，请求报文体积会更小。
性能：序列化和反序列化上，HTTP 大多是 JSON 数据格式，性能不如 gRPC 中的 protobuf；
还有基于 TCP 的 RPC 协议肯定能做到连接复用，但是 HTTP 协议就不一定了。

[应用场景]
RPC 主要侧重后端服务之间进行相互调用（也可以是移动端的接口调用），因为通过它进行后端服务间的调用，系统将具有很好的性能和拓展性。

[PS 常见通信协议]
网络接口层：...
网络层：IP、ICMP、ARP...
传输层：TCP、UDP
应用层：FTP、HTTP、HTTPS、SMTP、DNS...
（Socket 不是协议，是一种连接模式）

【实战】
rpc 的接口方法定义，一般是两个参数，一个作为 方法参数，一个作为 方法结果

测试发现：
- 第二参数一定是 * 类型（第一个参数则没有要求）
- 第一个参数并没有要求类型，即使是 * 类型，服务端对其做修改，客户端调用完成，值也不会发生改变

详见源码：

net/rpc/server.go:279：
	// suitableMethods returns suitable Rpc methods of typ, it will report
	// error using log if reportErr is true.

net/rpc/server.go:289：
	// Method needs three ins: receiver, *args, *reply.
	（方法有且仅有三个参数 0：receiver、1：*args、2：*reply）

net/rpc/server.go:296：
	// First arg need not be a pointer.
	（args 不要求指针类型）
	一定得是 builtin 类型或者是 exported 类型（结构体或者别名类型）

net/rpc/server.go:304
	// Second arg must be a pointer.
	（reply 一定得是指针类型）
	一定得是 builtin 类型或者是 exported 类型（结构体或者别名类型）

net/rpc/server.go:319
	// Method needs one out.
	（方法返回值 有且仅有一个）

net/rpc/server.go:326
	// The return type of the method must be error.
	（第一个返回值类型一定得是 error 类型）
*/