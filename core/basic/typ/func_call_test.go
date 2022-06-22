package typ

import (
	"testing"
)

func TestFuncCall(t *testing.T) {
	changeTest(t)
	anonymousTest(t)
}

// 普通方法调用
// 下面的语法糖，被称为 “方法值的正规化” 详见 evaluate_test.go
func normalTest(t *testing.T) {
	s.nor()
	// 实际：(*S).ptr(&s)（详见 define.go）
	// 这是语法糖，注意只对可寻址的值类型的属主有效（详见下面的 anonymousTest）
	// 重点：你会发现方法内部是否能够真是改变字段值，只取决于方法定义的 Receiver 是否带 *（详见下面的 changeTest）
	s.ptr()

	sPtr.nor() // 实际：S.nor(*sPtr)
	sPtr.ptr()

}

// 方法能够改变结构体字段值，只取决于方法的 Receiver 是否带星
func changeTest(t *testing.T) {
	// s.nor()
	// s.ptr() // ✅
	// (&s).nor()
	// (&s).ptr() // ✅

	// sPtr.nor()
	// sPtr.ptr()
	// (*sPtr).nor()
	// (*sPtr).ptr() // ✅

	// t.Logf("\"%s\"\n", s.Name)
	// t.Logf("\"%s\"\n", sPtr.Name)
}

// 匿名调用
func anonymousTest(t *testing.T) {
	S{}.nor()
	// S{}.ptr()  // 编译不通过
	(&S{}).nor()
	(&S{}).ptr()
}
