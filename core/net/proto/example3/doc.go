package example3

// 重点测试一下配置项 package、option go_package；以及 protoc 命令参数
//
// 为了后续指代方便和清晰，这里列一下语法：
// option go_package="<args1>"
// option go_package="<path>"
// option go_package="<args1>;<args2>"
// option go_package="<path>;<pack>"
//
// 测试记录：
// - option go_package 为必填参数
// - option go_package 只填一个代表目录时，最后一级作为实际包名
// - option go_package 填两个参数时，第一个作为生成目录、位置，第二个作为实际包名
// - option go_package 第二个参数中 “-” 为非法字符（单个参数则没有任何问题，目录也会正常生成）
//
// 命令参数 --go_out=plugins=grpc:<dir> 中的 dir 参数是要求必须存在的文件夹
//
// - 不同的 proto 文件， option go_package 的 args1 不能相同
//
// [引用]
// 能够被 import 的 proto 文件，其目录一定是在 File → Setting → Language & Frameworks → Protocol Buffers 的 Import Paths 中
// proto 中没有任何语法异常（特指 import）但是生成的 .pb.go 就不一定了，会直接将引用的目标 proto 的 go_package 的 args1 作为包名，
// Go 项目中的包引用都是相对于项目根目录来说的，所以会报错。
// TODO 暂未解决
//  - 到项目根文件所在目录执行生成命令（这...）
//  - 尝试调整 import → 失败
//  - 相关的包定义都用 github 开头的包名

// 同一个目录下：
//
// [pack1.proto]
// syntax = "proto3";
// option go_package = "./pack;a";
// message Model {string name = 1;}
//
// [pack2.proto]
// syntax = "proto3";
// option go_package = "./pack;a";
// message Model {string name = 1;}
//
// “protoc --go_out=plugins=grpc:. *.proto”
// 命令只会得到在处理第二个 proto 文件时报出关于 message 的 already define 的错误提示
//  → 添加 package p1; package p2; → 再次执行命令
//  → “protoc-gen-go: Go package "./pack" has inconsistent names a (pack1.proto) and b (pack2.proto)”
