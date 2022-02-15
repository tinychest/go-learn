package grpc

/*
【gRPC】

[参考]
https://chai2010.cn/advanced-go-programming-book/ch4-rpc/readme.html
https://mp.weixin.qq.com/s/5vbhRdqGAiQDFo_9YeiS2g

https://grpc.io/docs/what-is-grpc/introduction/
https://grpc.io/docs/what-is-grpc/core-concepts/

【概念】
[简介]
Like many RPC systems, gRPC is based around the idea of defining a service, specifying the methods that can be called remotely with their parameters and return types.
By default, gRPC uses protocol buffers as the Interface Definition Language (IDL) for describing both the service interface and the structure of the payload messages.
It is possible to use other alternatives if desired.

不说的那么细节、专业，gRPC 是 Google 开发的一个高性能开源的 RPC 框架，基于 HTTP/2，同时支持大多数流行的编程语言。

服务端响应信息：状态详细信息（status code and optional status message）、尾随元数据（optional trailing metadata）
客户端请求的参数信息：客户端元数据（client metadata）、方法名（method name）、deadline

[元数据（Metadata）]
元数据是有关特定 RPC 调用的信息（例如身份验证详细信息）
采用键值对列表的形式，键是 string 类型，值一般也是 string 类型，也可以是二进制数据。
元数据对于 gRPC 本身来说，是不透明的 - 它允许客户端提供与服务端调用相关的信息，反之亦然。

对元数据的访问取决于语言。

[通道（Channels）]
> 这个原文就很简明易懂了，不翻译了
A gRPC channel provides a connection to a gRPC server on a specified host and port.
It is used when creating a client stub. Clients can specify channel arguments to modify gRPC’s default behavior,
such as switching message compression on or off. A channel has state, including connected and idle.

How gRPC deals with closing a channel is language dependent. Some languages also permit querying channel state.

[服务定义（Service definition）]
就像很多 RPC 系统，gRPC 是基于定义服务的思想，指定可以被远程调用的方法极其参数和返回类型

默认情况下，gRPC 使用 protocol buffers 作为 IDL 去描述 service interface 和 structure of the payload messages，
如果需要的话，也可以使用其他替代方案

【生命周期】
gRPC 不像传统的 RPC 方法调用，不仅支持 一请求一响应；还拓展了流式的 一请求多响应 和 多请求多响应。

[一元 RPC（Unary RPC）]
Unary RPCs where the client sends a single request to the server and gets a single response back, just like a normal function call.

发起一次请求，给予一次响应

[服务流 RPC（Server stream RPC）]
Server streaming RPCs where the client sends a request to the server and gets a stream to read a sequence of messages back.
The client reads from the returned stream until there are no more messages. gRPC guarantees message ordering within an individual RPC call.

除了服务端返回一个消息流来响应客户端的请求，其他和一元 RPC 一致

[客户端 RPC（Client stream RPC）]
Client streaming RPCs where the client writes a sequence of messages and sends them to the server, again using a provided stream.
Once the client has finished writing the messages, it waits for the server to read them and return its response.
Again gRPC guarantees message ordering within an individual RPC call.

除了客户端发送消息流给服务端，其他和一元 RPC 一致

[双向流 RPC（Bidirectional stream RPC）]
Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream.
The two streams operate independently, so clients and servers can read and write in whatever order they like:
  for example, the server could wait to receive all the client messages before writing its responses,
  or it could alternately read a message then write a message, or some other combination of reads and writes.
The order of messages in each stream is preserved.

调用由客户端发起调用时，进行初始化，服务端会收到客户端的 meta data、method name、deadline。
服务端可以返回自己的初始化元素据也可以等待客户端发起数据流
客户端流和服务端流是相互独立的，所以客户端和服务端可以任意顺序读写消息；举例来说，服务端可以等待接受完客户端的所有消息后，再进行响应；
客户端和服务端之间也可以进行一问一答 - 服务端接收到一个请求，然后给予响应，客户端可以基于这个响应发起另外一个请求，以此类推

[Deadlines/Timeouts]
gRPC 允许客户端指定在 RPC 因 DEADLINE_EXCEEDED 错误而终止之前，他们愿意等待 RPC 完成的时间。
这也是视语言而定的，有些根据超时的 duration，有些根据超时时间点，有些也可能没有默认期限。

[RPC termination]
在 RPC 中，客户端和服务端对于调用是否成功都有他们各自本地、独立的判断；也就是说，它们的结论可能不一致，
例如，服务端发送完了所有响应并认为 RPC 成功并完成，但是客户端收到时都超时了；还有如服务端在客户端发送所有请求之前的某个时间点就认为 RPC 调用完成了。

[Cancelling RPC]
无论是客户端还是服务端都可以随时取消 RPC。
取消操作会立即中值 RPC，以便不再进行任何工作。

【核心概念】
[IDL（Interface Description Language）]
接口描述语言，例如 proto buffer。一些跨平台的 RPC 框架可以根据 IDL，在编译时期使用代码生成器生成 stub 代码。

[Protocol Buffer]
详见，同级目录的 proto 包

【实战】
[Using the API]
首先从在 .proto 文件中定义一个 service 开始，gRPC 提供了生成 client-and-server-side code 的 protocol buffer compiler plugins。
gRPC 使用者通常在客户端调用这些 API，并在服务端实现相应的 API。
- On the server size, the server implements the methods declared by service and runs a gRPC server to handle client calls.
  The gRPC infrastructure decodes incoming requests, executes service methods, and encodes service responses.

> local object 本地对象
- On the client side, the client has a local object known as stub (for some languages, the preferred term is client)
  that implements the same methods as the service. The client can then just call those methods on the local object,
  wrapping the parameters for the call in the appropriate “protocol buffer” message type - gRPC looks after（负责）
  sending the request(s) to the server and returning the server's “protocol buffer” response(s).

【其他】
[HTTP/2]
> 特别提及的基础协议 HTTP/2
HTTP/2 采用二进制格式传输协议，HTTP/1.x 是文本格式。
HTTP/2 支持通过一个连接发送多个并发请求。
HTTP/2 中，服务器可以对客户端的一个请求给予多个响应（服务端推送）；这是 HTTP/1.x 做不到的。
HTTP/2 对消息头进行了压缩，能够节省网络流量。

https://pkg.go.dev/encoding/gob
*/
