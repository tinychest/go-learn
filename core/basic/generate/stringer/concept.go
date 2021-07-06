package stringer

/*
	1.定义一个基础类型为 int 的类型作为枚举，同时准备好具体的枚举示例
	2.准备：go get -u golang.org/x/tools/cmd/stringer
	3.执行: stringer -type=Pill
	4.生成一个名为 pill_string.go 的源码文件，里边为 Pill 实现了 string 方法
	注意：文件名不能以 test 结尾，会提示找不到 stringer: no values defined for type Pill

	不添加指定的注释，这样也行
    1.在枚举类型上添加指定的 go generate stringer 注释
    2.执行：go generate

	其他：生成的 xxx_string.go 会被 goland 归到 xxx.go 的下边
*/