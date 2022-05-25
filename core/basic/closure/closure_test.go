package closure

import "testing"

// 闭包函数是什么，闭包函数就是有着变量捕获列表的函数
// - 闭包很容易造成引用的变量发生逃逸
//
// 下面的若干的测试样例都是为了说明闭包函数对引用变量，引用的是其地址，所以能够改变其值。

func TestClosure(t *testing.T) {
	outModifyTest(t)
	inModifyTest(t)
	inPassModifyTest(t)
}

func outModifyTest(t *testing.T) {
	n := 0
	f := func() { t.Log(n) }
	n++

	f()
}

func inModifyTest(t *testing.T) {
	n := 0
	f := func() { n++ }

	f()
	t.Log(n)
}

func inPassModifyTest(t *testing.T) {
	n := 0
	f := func() { n++ }

	whatever(f)
	t.Log(n)
}

func whatever(f func()) {
	f()
}
