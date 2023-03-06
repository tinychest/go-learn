package _defer

import "testing"

// 下面两组例子想阐述的 defer 原理是一样的
func TestDeferReturn(t *testing.T) {
	t.Log(returnTest1())
	t.Log(returnTest2())
}

func returnTest1() int {
	var n int
	defer add(&n)
	// 确定返回值 0 → 执行 defer：n++
	return n
}

func returnTest2() (n int) {
	defer add(&n)
	// 确定返回值 n → 执行 defer：n++
	return
}

func add(i *int) {
	*i++
}
