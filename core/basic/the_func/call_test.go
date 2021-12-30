package the_func

import (
	"testing"
)

func TestFuncCall(t *testing.T) {
	// special()
	change()
}

// 普通方法调用
// 下面的语法糖，被称为 “方法值的正规化” 详见 evaluate_test.go
func normal() {
	s.nor()
	s.ptr() // 这是语法糖（此语法糖只对可寻址的值类型的属主有效），编译会隐式改写为 (&s).ptr()（在 C 语言中是 (&s)->ptr()）

	sPtr.nor() // 同上，最终改写为 (*sPtr).nor()
	sPtr.ptr()
}

// 方法能够改变结构体字段值，只取决于方法的 Receiver 是否带星
func change() {
	// s.nor()
	// s.ptr() // ✅
	// (&s).nor()
	// (&s).ptr() // ✅

	// sPtr.nor()
	// sPtr.ptr()
	// (*sPtr).nor()
	// (*sPtr).ptr() // ✅

	// fmt.Printf("\"%s\"\n", s.Name)
	// fmt.Printf("\"%s\"\n", sPtr.Name)
}

// 匿名调用
func anonymous() {
	S{}.nor()
	// S{}.ptr()  // 编译不通过
	(&S{}).nor()
	(&S{}).ptr()
}

// 特殊调用形式（详见 define.go 了解为什么）
func special() {
	S.nor(s)
	// S.ptr(s) // Goland 不提示，实际运行 invalid method expression S.ptr (needs pointer receiver: (*S).ptr)
	// S.nor(sPtr) // 编译不通过
	// S.ptr(sPtr) // 编译不通过

	// (*S).nor(s) // 编译不通过
	// (*S).ptr(s) // 编译不通过
	(*S).nor(sPtr)
	(*S).ptr(sPtr)
}