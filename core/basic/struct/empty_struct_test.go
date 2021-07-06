package _struct

import (
	"fmt"
	"go-learn/util"
	"testing"
)

type tStruct struct{}

func (t *tStruct) Func() {
	fmt.Println("call...")
}

// 空结构体实例类型是不占用空间的 unsafe.Sizeof(struct{}{})
func TestEmptyStruct(t *testing.T) {
	// 例1：set - Go 中没有提供 set 的数据类型，只能用 map 来模拟替代
	var set = util.NewSet(4)
	set.Add("123")

	// 例2：信号通道 - 我们在日常业务逻辑使用通道的时候，经常会遇到就仅仅发送信号，有可能发送一个信号，通道就没用了，或者程序就结束了
	anonymousSignal := make(chan struct{})
	anonymousSignal <- struct{}{}

	// 例3：nil 结构体实例也能调用为其定义绑定的方法
	// 仅提供方法的结构体，方法的 Receiver 选用 int bool 类型都会占用空间
	// 空指针异常（×） 能够调用（√）
	var tPtr *tStruct
	tPtr.Func()
}
