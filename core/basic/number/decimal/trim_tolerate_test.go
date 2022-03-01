package decimal

import (
	"testing"
)

func TestBadAndTolerate(t *testing.T) {
	// [Ceil]
	// big := 0.07 * 100 // 实际是 7.00...001
	t.Log(BadCeil(0.07, 2))
	t.Log(ToleCeil(0.07, 2))

	// [Floor]
	// 换成 const 定义，结果是正确的（实际上常量值得出的是小的误差值的案例不好找）
	var a, b = 0.1, 0.7
	var args = a + b
	t.Log(BadFloor(args, 1))
	t.Log(ToleFloor(args, 1))

	// [Round]
	a, b = 0.03, 0.005
	args = a + b
	t.Log(BadRound(args, 2))
	t.Log(ToleRound(args, 2))
}
