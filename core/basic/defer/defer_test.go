package _defer

import "testing"

// 执行顺序：语句由上至下执行 → 确定 return 返回值 → 由下至上执行 defer 的函数

func TestDefer(t *testing.T) {
	defer t.Log(1)
	defer t.Log(2)
	t.Log(3)
}
