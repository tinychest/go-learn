package todo

import (
	"go-learn/util"
	"testing"
)

// 测试个屁，值类型在栈内存中（调用一个方法，就会开辟一个方法栈）
func TestInterfaceWrapAddr(t *testing.T) {
	type p struct{}

	// 测试
	var p1 interface{} = p{}
	var p2 = p1
	util.PrintAddr(p1)
	util.PrintAddr(p2)
}
