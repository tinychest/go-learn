package example5

/*
参照 gin 的 middleware 概念，gRPC 中也可以为所有 RPC 方法定义统一的前置处理逻辑

[全局]
- 如何注册中间件

1.发现方法名相关的方法（Unary、Interceptor）
func ChainUnaryInterceptor(interceptors ...UnaryServerInterceptor) ServerOption

2.如何使用：上面方法的返回值正好对应了 NewServer 方法的参数类型
func NewServer(opt ...ServerOption) *Server

3.如何定义：UnaryServerInterceptor 肯定是个接口类型，定义实现就好

[设置权限标识]
类比 HTTP 的 Request Header

类比 HTTP Server 中的 Request 进行类推，发现 xxx 中的 req 中并没有相关参数信息，
而在 context 中发现了默认的请求头参数，由此推断 context 是重心.

但是调试发现在 Client 直接给 context 赋值，在 Server 端的 context 中根本没有相关数据

尝试在源码中索引出相关方法：
- SetHeader(ctx context.Context, md metadata.MD) error
  SendHeader(ctx context.Context, md metadata.MD) error
  都是 Header 相关的方法，但是实际是 stream 的 Header 发送方法，在 Unary 环境中使用，会得到：
  failed to fetch the stream from the context context.Background
  failed to fetch the stream from the context context.Background.WithValue(type metadata.mdOutgoingKey, val <not Stringer>)
- 通过插件生成的 Client 的实现代码：Hello(ctx context.Context, in *HelloArgs, opts ...grpc.CallOption) (*HelloReply, error)
  方法的第三个参数类型 和 Header(md *metadata.MD) CallOption 方法返回值类型对上了
  （实际发现，这是 Client 调用结束后的后置处理，用于重置 context 中的头；目前刚刚接触这一块，自然不懂，后面会知道的）

查找资料发现应该通过 metadata 包：
客户端通过 metadata.NewOutgoingContext 进行请求头设置
服务端通过 metadata.FromIncomingContext 获取请求头

其实这里就是官方概念文档中 元数据的概念

[其他]
stream 定义 Interceptor 的方式和 Unary 都有不同、设置请求头的方式 方法

TODO 应当找到官方文档样例，再来敲的，这不是一个正确的学习方式
*/
