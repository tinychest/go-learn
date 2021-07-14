package _defer

import (
	"fmt"
	"testing"
)

// TestDefer 下面两组例子想阐述的 defer 原理是一样的
func TestDefer(t *testing.T) {
	funcTest1()
	funcTest2()

	println(returnTest1())
	println(returnTest2())
}

func funcTest1() {
	var n = 0

	defer func(a int) {
		fmt.Println(a)
	}(n) // 此时就已经固定了 n 的参数值为 0

	n = 1
}

func funcTest2() {
	var n = 0

	// 闭包
	defer func() {
		fmt.Println(n)
	}()

	n = 1
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
