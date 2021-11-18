package string

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
【引出】
原文：https://mp.weixin.qq.com/s/jztwFH6thFdcySzowXOH_Q
问：string ←→ []byte 会发生内存拷贝么？
答：会

【原因】
在 Go 中 string 类型被设计为不可变的
这样在并发场景下，我们可以在不加锁的控制下， 多次使用同一字符串，在保证高效共享的情况下而不用担心安全问题

【基础】
查看 byte 的定义：type byte = uint8
查看 string 的定义：type string string

string → []byte
string → []rune

【分析】
可以通过命令：go tool compile -N -l -S ./string_to_byte/string.go 定位到调用的底层方法
可以通过定位到源码的 stringtoslicebyte、stringbytetostring 查看实现细节

通过查看源码，发现两个方法的标准转换都发生了内存拷贝，实际上可以通过强转换来避免内存拷贝和内存申请
详见提供的实现：slicebytetostringtmp、stringtoslicebytetmp

【其他】
Go 语法提供的类型转换和高效的底层直接类型转换（避免拷贝，但是不安全，甚至 panic 都 recover 不了）
string 类型虽然是不能更改的，但是可以被替换，因为 stringStruct 中的 str 指针是可以改变的，只是指针指向的内容是不可以改变的
*/

func TestStrAddr(t *testing.T) {
	// 【safe】
	// 可以看出来，指针指向的位置发生了变化，也就说每一个更改字符串，就需要重新分配一次内存，之前分配的空间会被回收
	s := `abc`
	bs1 := []byte(s)
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", bs1)

	// 【unsafe】
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	bs2 := *(*[]byte)(unsafe.Pointer(&ret))
	fmt.Printf("%p\n", bs2)

	// 没问题
	s = "qqq"

	// gg，recover 都救不回来
	bs2[0] = 1
}

// TODO 如果不发生内存拷贝，打印的地址不同？
// TODO 这是 Go 底层的方法，就算真的能保证不发生拷贝，但是调用方法的时候，不就已经发生了拷贝？
func stringtoslicebytetmp(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	return *(*[]byte)(unsafe.Pointer(&ret))
}
