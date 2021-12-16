package _struct

import (
	"testing"
)

type out struct {
	OutName string
	in
}

type in struct {
	InName string
}

func (i in) f() {}

func TestAnonymousStruct(t *testing.T) {
	// 例1 外部类实例能够直接读写内部类的字段；不能结构体写，编译不通过
	var o out
	o.InName = "123"
	println(out{}.InName)

	// 例2 外部类实例能够直接调用内部类的方法
	o.f()

	// 例3 in 类型不能接受 out{in}，反之亦然
	// var _ in = out{}
	// var _ out = in{}
}
