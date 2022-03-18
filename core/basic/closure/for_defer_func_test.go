package closure

import (
	"testing"
)

var s = []string{"a", "b", "c"}

func TestDeferFunc(t *testing.T) {
	mistake(t)
	tryFix(t)
	fix1(t)
	fix2(t)
}

func mistake(t *testing.T) {
	for _, v := range s {
		defer func() {
			t.Log(v)
		}()
	}
}

// 失败原因：并未改变函数内对 for 局部变量的闭包引用
func tryFix(t *testing.T) {
	for _, v := range s {
		defer func() {
			i := v
			t.Log(i)
		}()
	}
}

// 局部函数对未定义参数 - 外部变量的直接引用称作闭包
func fix1(t *testing.T) {
	for _, v := range s {
		v := v
		defer func() {
			t.Log(v)
		}()
	}
}

// 为什么能解决：不要被冗杂的语法迷惑了双眼，循环体只要看成一个函数的调用即可
// 方法传参是对变量值的引用，闭包是对变量的引用
func fix2(t *testing.T) {
	for _, v := range s {
		// defer func(v string) {
		// 	t.Log(v)
		// }(v)
		defer t.Log(v)
	}
}
