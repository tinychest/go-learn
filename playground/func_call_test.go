package playground

import (
	"fmt"
	"testing"
)

/* interface、func call playground */
type Killer interface {
	Kill()
}

type Jack struct{}

func (j Jack) Kill() {
	fmt.Println("slash slash slash")
}

func TestMess(t *testing.T) {
	// 通过实例直接调用方法
	Jack{}.Kill()
	// 直接调用编译器生成的对应方法
	Jack.Kill(Jack{})
	// 通过匿名接口的形式调用目标实例的方法
	interface{ Kill() }.Kill(Jack{})
	// 将实例通过接口类型断言后进行匿名接口断言，再调用匿名接口定义的方法
	((Killer)(Jack{})).(interface{ Kill() }).Kill()
}
