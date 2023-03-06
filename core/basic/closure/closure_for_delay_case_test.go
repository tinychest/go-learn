package closure

import (
	"testing"
)

// closure：闭包
// for：在 for 循环中
// delay： defer or go 关键字都能起到异步，延迟执行的效果
//
// 以一个最简单的闭包问题为背景，循序渐进，演示解决步骤，同时阐明闭包问题的核心

var arr = []int{1, 2, 3}

func TestDeferFunc(t *testing.T) {
	// mistake(t)
	// tryFix(t)
	// fix1(t)
	// fix2(t)
	// extend3(t)
}

// 闭包问题，没有按照预期打印
func mistake(t *testing.T) {
	for _, v := range arr {
		defer func() {
			t.Log(v)
		}()
	}
}

// 错误。仍没有按照预期打印
func tryFix(t *testing.T) {
	for _, v := range arr {
		defer func() {
			i := v
			t.Log(i)
		}()
	}
}

// 正确。解决关键在于每个函数绑定的闭包变量都是一个新的
func fix1(t *testing.T) {
	for _, v := range arr {
		v := v
		defer func() {
			t.Log(v)
		}()
	}
}

// 正确。解决关键在于每个函数绑定的闭包变量都是一个新的
func fix2(t *testing.T) {
	for _, v := range arr {
		defer func(v int) {
			t.Log(v)
		}(v)
	}
}

// 闭包的一个关键是局部函数的定义，而下面没有，所以不存在闭包问题
func extend1(t *testing.T) {
	for _, v := range arr {
		defer t.Log(v)
	}
}

// 拓展说明一下，for 循环背景下，闭包问题的一个要素，就是延时
func extend2(t *testing.T) {
	// 没有问题
	for _, v := range arr {
		func() {
			t.Log(v)
		}()
	}
	// 延时1
	for _, v := range arr {
		defer func() {
			t.Log(v)
		}()
	}
	// 延时2
	fs := make([]func(), 0, len(arr))
	for _, v := range arr {
		fs = append(fs, func() {
			t.Log(v)
		})
	}
	for _, f := range fs {
		f()
	}
}
