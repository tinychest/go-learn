package closure

import "testing"

// 闭包函数就是有着变量捕获列表的函数（捕获的变量会逃逸到堆上）
// 捕获变量本质是对地址的引用，所以能够更变其值

func TestClosure(t *testing.T) {
	// 这两个 case 没有本质区别
	case1Test(t)
	case2Test(t)
}

// 闭包函数外对捕获变量的值进行修改
func case1Test(t *testing.T) {
	n := 0
	f := func() { t.Log(n) }
	n++

	f()
}

// 闭包函数外对捕获变量的值进行修改
func case2Test(t *testing.T) {
	n := 0
	f := func() { n++ }

	f()
	t.Log(n)
}
