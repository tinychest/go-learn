package _struct

import (
	"go-learn/util"
	"testing"
)

type out struct {
	OutName string
	in
}

type in struct {
	InName string
}

func TestAnonymousStruct(t *testing.T) {
	// 编译不通过，不能直接给 name2 属性赋值
	// _ := out {InName: "2"}

	// 编译通过，可以直接调用属性
	util.Use(out{}.InName)
}
