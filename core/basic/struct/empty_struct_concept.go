package _struct

import (
	"fmt"
	"go-learn/tool/set"
)

// 空结构体实例类型是不占用空间的 unsafe.Sizeof(struct{}{})

func addrConcept() {
	var a, b struct{}
	// true
	fmt.Println(a == b)
	// addr equal
	print(&a, "\n", &b, "\n")
	// false
	fmt.Println(&a == &b)
}

type tStruct struct{}

func (t *tStruct) Func() {
	fmt.Println("call...")
}

func useConcept() {
	// 例1：set - Go 中没有提供 set 的数据类型，只能用 map 来模拟（map[string]struct{}）
	var s = set.New[string](4)
	s.Add("123")

	// 例2：信号通道 - 我们在日常业务逻辑使用通道的时候，经常会遇到通道传递的数据本身不需要有任何意义，该通道在整个生命周期只接收一个信号，被读取后就无用了，或者程序就结束了
	sig := make(chan struct{})
	sig <- struct{}{}

	// 例3：nil 结构体实例也能调用为其定义绑定的方法
	// 仅提供方法的结构体，方法的 Receiver 选用 int bool 类型都会占用空间
	// 空指针 panic（×） 正常调用（√）
	var tPtr *tStruct
	tPtr.Func()
}
