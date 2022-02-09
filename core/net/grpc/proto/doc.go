package pro

/*
【Protocol Buffer】
Protobuf 是 Protocol Buffers 的简称，是一种和语言、平台无关，可拓展的序列化、结构化数据的数据描述语言
就像 JSON 一样，作为一种通用的数据格式规范，谷歌为此定义，用作 RPC 数据传输的规范

【基本语法】
详见 hello.proto 的注释

【安装系统编译工具】
https://github.com/protocolbuffers/protobuf
protoc-3.19.1-win64.zip

【安装工具】
go get github.com/golang/protobuf/protoc-gen-go（命令过时）
go get -d github.com/golang/protobuf/protoc-gen-go（依赖过时）

go get -d google.golang.org/protobuf
或
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

【.proto → .go】
protoc --go_out=. hello.proto
protoc --go_out=plugins=grpc:. hello.proto
*/
