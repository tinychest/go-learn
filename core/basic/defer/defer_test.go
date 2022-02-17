package _defer

import (
	"testing"
)

// TestDefer 下面两组例子想阐述的 defer 原理是一样的
func TestDefer(t *testing.T) {
	t.Log(returnTest1())
	t.Log(returnTest2())
}

func returnTest1() int {
	var n int
	defer func() {
		n++
	}()
	// 确定返回值 0 → n++ → 返回 0
	return n
}

func returnTest2() (n int) {
	defer func() {
		n++
	}()
	// 确定返回值 n → n++ → 返回 n
	return
}
