package example1

/*
定义一个较为完整的 proto、基于 grpc 的包 给出样例
- 编写 proto（包含简单方法、数据实体定义）
- 生成代码（因为定义了方法，所以一定要通过 gRPC 插件生成）
- 根据 生成的代码 和 grpc 包的相关方法提示，实现 客户端 和 服务端
- 运行服务端、运行客户端

（grpc 中的方法签名似乎没有 rpc 中定义的规则限制）

【使用 Postman 调用定义的 gRPC 服务端实现】
[APIs] define proto
➕（Create new API）
	Name：hello
	Version：1.0.0
	Scheme type：Protobuf 3
点击 Create API → 点记 Close → 出现 hello/1.0.0

点击 1.0.0 → Definition → 复制粘贴 hello.proto 的内容 → Save

[New] gRPC Request
Workspace → New → gRPC Request
	Enter Server URL：127.0.0.1:1234
	Choose...：hello
	Select a method：Hello

Message → 点击 Generate Example Message 按钮

点击 Invoke 按钮，成功调用
（因为还处于测试阶段，所以 gRPC Request 创建的实例还无法保存）
*/
