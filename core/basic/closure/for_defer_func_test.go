package closure

import (
	"testing"
)

// 以一个最简单的闭包问题为背景，循序渐进，演示解决步骤，同时阐明闭包问题的核心

var x = []int{1, 2, 3}

func TestDeferFunc(t *testing.T) {
	// mistake(t)
	// tryFix(t)
	// fix1(t)
	// fix2(t)
	// extend3(t)
}

// 1.没有按照预期打印
func mistake(t *testing.T) {
	for _, v := range x {
		defer func() {
			t.Log(v)
		}()
	}
}

// 2.仍没有按照预期打印
func tryFix(t *testing.T) {
	for _, v := range x {
		defer func() {
			i := v
			t.Log(i)
		}()
	}
}

// 3.正确。解决关键在于每个函数绑定的闭包变量都是一个新的
func fix1(t *testing.T) {
	for _, v := range x {
		v := v
		defer func() {
			t.Log(v)
		}()
	}
}

// 4.闭包的关键一个是局部函数，一个是函数内部对上下文变量的引用；很明显下面没有定义局部函数
func fix2(t *testing.T) {
	for _, v := range x {
		defer t.Log(v)
	}
}

// 5.拓展说明一下，for 循环背景下，闭包问题的一个要素，就是延时
func extend3(t *testing.T) {
	// 没有问题
	for _, v := range x {
		func() {
			t.Log(v)
		}()
	}
	// 延时1
	for _, v := range x {
		defer func() {
			t.Log(v)
		}()
	}
	// 延时2
	fs := make([]func(), 0, len(x))
	for _, v := range x {
		fs = append(fs, func() {
			t.Log(v)
		})
	}
	for _, f := range fs {
		f()
	}
}
