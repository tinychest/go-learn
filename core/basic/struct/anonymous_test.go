package _struct

import (
	"testing"
)

func TestAnonymous(t *testing.T) {
	// anonymousStructTest(t)
	specialTest(t)
}

func anonymousStructTest(t *testing.T) {
	var s1 *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}
	var s2 *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}
	// the underlying types are identical
	// s1 和 s2 是相同的类型，这里主要想强调 tag 必须也是相同的才算相同的类型（很合理，不然在某些方面的表现出现差异，也就不能称之为相同了）
	t.Log(s1 == s2)
}

// 下面样例，实际不会 panic
// 源自：https://mp.weixin.qq.com/s/s5JoyPbBzhu_GjGDd_pJYw
// 核心在于 len 函数
//    返回结果总是 int；
//    返回结果有可能是常量；
//    有时对函数参数不求值，即编译期确定返回值；
//
// 如果 len 或 cap 的函数参数 v 是字符串常量，则返回值是一个常量。
//
// 如果 v 的类型是数组或指向数组的指针，且表达式 v 没有包含 channel 接收或（非常量）函数调用，则返回值也是一个常量。
// 这种情况下，不会对 v 进行求值（即编译期就能确定）。否则返回值不是常量，且会对 v 进行求值（即得运行时确定）。
func specialTest(t *testing.T) {
	// panic ss 没有初始化，为 nil（报的不是空指针，而是数组越界）
	// var ss []string
	// t.Log(ss[1])

	var x *struct {
		s [][32]byte
	}

	t.Logf("T = %T V = %v\n", x, x)
	t.Log(len(x.s[99]))
}
