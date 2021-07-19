package slice

import (
	"fmt"
	"testing"
)

// 参考：https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651448705&idx=3&sn=dd19cc4218dfd5956d3e70f8a760868c&scene=21#wechat_redirect
// 注意：Go 语法提供的类型转换和高效的底层直接类型转换（避免拷贝，但是不安全，甚至 panic 都 recover 不了）

// 分析流程
// 查看 byte 的定义：type byte = uint8
// 查看 string 的定义：type string string

// 通过类型定义和构建方法的源码：string 的本质就是 []byte

// 问：string 类型为什么还要再数据的基础上在进行一次封装呢？
// 在 Go 中 string 类型被设计为不可变的（其他语言中 string 类型也是被设计为不可变的）
// 这样好处是，在并发场景下，我们可以在不加锁的控制下， 多次使用同一字符串，在保证高效共享的情况下而不用担心安全问题

// TestStrAddr string 类型虽然是不能更改的，但是可以被替换，因为 stringStruct 中的 str 指针是可以改变的，只是指针指向的内容是不可以改变的
func TestStrAddr(t *testing.T) {
	// 可以看出来，指针指向的位置发生了变化，也就说每一个更改字符串，就需要重新分配一次内存，之前分配的空间会被gc回收
	str := `abc`
	fmt.Printf("%p\n", []byte(str))
	str = `def`
	fmt.Printf("%p\n", []byte(str))
}

// 可以通过命令：go tool compile -N -l -S ./string_to_byte/string.go 定位到调用的底层方法
// 可以通过定位到源码的 stringtoslicebyte、stringbytetostring 查看实现细节

// 通过查看源码，发现两个方法的标准转换都发生了内存拷贝，实际上可以通过强转换来避免内存拷贝和内存申请
// 详见提供的实现：slicebytetostringtmp、stringtoslicebytetmp

// 结论
// 如果是在高性能场景下使用，是可以考虑使用强转换的方式的，但是要注意强转换的使用方式，它不是安全的
// string → []byte 后，如果尝试修改字符，那么会出现 defer + recover 都无法恢复的错误
