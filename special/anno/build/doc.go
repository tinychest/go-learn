package build

/*
【引出】
当前包下的 a.go 和 b.go 中都定义了 hello 方法，但是编译时，编译器并没有报出错误

【解释】
这就是基于标签的条件编译，借由 +build <tag1> <tag2> ... 这样的注释语法

只要编译命令时指定了标签，就编译或者不编译，如：
go build -tags hello（Goland：File → Settings... → Go → Build Tags & Vendoring/Custom Tags）
// +build hello（编译）
// +build !hello（不编译）

【注意】
这样的注释必须在包名上方，且必须有一段空行（否则会作为包注释）

【其他】
条件编译的条件不限于 tag，还有文件命名

以 _linux.go 结尾的文件只会在 linux 系统下编译
以 _windows_amd64.go 结尾的文件只会在 windows 64bit 系统下编译
还有如：_freebsd_arm.go、_plan9.go

TODO go:build？
 */
