package example6

// 流式客户端和流式服务端交互的中间参杂请求头的发送和接收
// 1.将 example 完整复制过来，修改每个文件中的包名
// 2.每个数据发送、接收的前面发送、接收请求头

// - 发现 grpc 返回错误的正确姿势

// TODO [可以补充的概念]
//   各种特殊的类型需要导入 google 的包才能使用，如（import "google/protobuf/any.proto"; import "hit-rpc/proto/base.proto";）
//   metadata 的概念
//   package 的概念，package 是 proto 中的概念，proto 之间的互相引用也是通过这个，和生成的 go 文件位置，包名都没关系

// TODO [待补充]
//   protoc 命令的参数 "protoc-go-inject-tag -XXX_skip=form"
//   在字段加上 "// @inject_tag: form:"id""
//   封装 UnaryEchoServer、StreamEchoServer（grpc 包下的 test 包，好像有类似的实现）
