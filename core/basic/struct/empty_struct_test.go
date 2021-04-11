package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

// 《空结构体实例的应用场景》
// 应用1：使用 map 来模拟 set - 例1
// 应用2：仅提供方法的结构体 - 方法的 Receiver 选用 int bool 类型都会占用空间
// 应用3：信号通道 - 例2
func TestEmptyStruct(t *testing.T) {
	// 空结构体实例类型是不占空间的
	println(unsafe.Sizeof(struct{}{}))

	// 例1：set - Go 中没有提供 set 的数据类型，只能用 map 来模拟替代
	type Set map[string]struct{}
	var set Set
	add := func(s string) {
		set[s] = struct{}{}
	}
	add("123")

	// 例2：信号通道 - 我们在日常业务逻辑使用通道的时候，经常会遇到就仅仅发送信号，有可能发送一个信号，通道就没用了，或者程序就结束了
	anonymousSignal := make(chan struct{})
	anonymousSignal <- struct{}{}
}

type tStruct struct{}

func (t *tStruct) Func() {
	fmt.Println("call...")
}

// 为 nil 的结构体指针类型实例，调用对应类型的方法
func TestNilStructMethodCall(t *testing.T) {
	var tPtr *tStruct
	// 空指针异常（×）
	// 能够调用（√）
	tPtr.Func()
}
