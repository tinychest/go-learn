package _defer

import "testing"

/*
将定义的 defer 语句放入栈底，还有 defer 则继续放在栈底
defer 的执行时机是在 return 返回值赋值后，执行 RET 是最终将返回值返回
*/

func TestOrder(t *testing.T) {
	orderTest1()
	orderTest2()
}

// 输出：132，需要好好反省
func orderTest1() {
	var s S
	defer s.f(1).f(2)
	print(3)
}

func orderTest2() {
	// 后执行
	defer func() {
		println(1)
	}()

	// 先执行
	defer func() {
		println(2)
	}()
}

// Prepare ---
type S struct{}

func (s S) f(n int) S {
	print(n)
	return s
}
