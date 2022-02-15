package pro

/*
【Protocol Buffer】
protocol 直译 协议
buffer 直译 缓冲

作为 gRPC 的核心之一

Protobuf 是 Protocol Buffers 的简称，是一种和语言、平台无关，可拓展的序列化、结构化数据的数据描述语言
就像 JSON、XML 一样，作为一种通用的数据格式规范；但是它编解码速度更快，数据传输体积更小

[基本语法]
详见 hello.proto 的注释

[安装系统编译工具]
https://github.com/protocolbuffers/protobuf
protoc-3.19.1-win64.zip

[安装工具]
go get github.com/golang/protobuf/protoc-gen-go（命令过时）
go get -d github.com/golang/protobuf/protoc-gen-go（依赖过时）

go get -d google.golang.org/protobuf
或
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

[.proto → .go]
protoc --go_out=. hello.proto
protoc --go_out=plugins=grpc:. hello.proto

第一次使用下面的命令，为含有 service 定义的 proto 文件生成代码，然后，需要执行 go mod tidy 来保证下面包的依赖引入：
"google.golang.org/grpc"
"google.golang.org/grpc/codes"
"google.golang.org/grpc/status"
*/
