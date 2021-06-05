package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

// 一、空结构体类型作为 map 的值类型（本质就是兼容添加重复元素的 Set）
type Set map[string]struct{}

func (s Set) Add(str string) {
	s[str] = struct{}{}
}

func (s Set) Exists(key string) (exists bool) {
	// 还可以有判断是否有指定元素的方法
	_, exists = s[key]
	return
}

func NewSet(sizes ...int) Set {
	var size int
	if len(sizes) > 0 {
		size = sizes[0]
	}
	return make(Set, size)
}

// 二、空结构体类型作为 Receiver 的类型
type tStruct struct{}

func (t *tStruct) Func() {
	fmt.Println("call...")
}

// 《空结构体实例的应用场景》
//（空结构体实例类型是不占用空间的）
// 应用1：使用 map 来模拟 set - 例1
// 应用2：信号通道 - 例2
// 应用3：仅提供方法的结构体，方法的 Receiver 选用 int bool 类型都会占用空间 - 例3
func TestEmptyStruct(t *testing.T) {
	// 空结构体实例类型是不占空间的
	println(unsafe.Sizeof(struct{}{}))

	// 例1：set - Go 中没有提供 set 的数据类型，只能用 map 来模拟替代
	var set = NewSet(4)
	set.Add("123")

	// 例2：信号通道 - 我们在日常业务逻辑使用通道的时候，经常会遇到就仅仅发送信号，有可能发送一个信号，通道就没用了，或者程序就结束了
	anonymousSignal := make(chan struct{})
	anonymousSignal <- struct{}{}
}

// 《空的结构体实例也能调用为其定义绑定的方法》
func TestNilStructMethodCall(t *testing.T) {
	var tPtr *tStruct
	// 空指针异常（×） 能够调用（√）
	tPtr.Func()
}