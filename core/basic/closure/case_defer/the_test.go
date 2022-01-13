package case_defer

import (
	"go-learn/core"
	"testing"
)

var ps = []core.Person{
	{Name: "小明"}, {Name: "小红"}, {Name: "小光"},
}

func TestDefer(*testing.T) {
	// mistake(t)
	// tryFix1(t)
	// tryFix2(t)
	// fix1(t)
	// fix2(t)
	fix3(t)
}

func mistake(t *testing.T) {
	for _, item := range ps {
		defer func() {
			t.Log(item.Name)
		}()
	}
}

// 失败原因：并未改变函数内对 for 局部变量的闭包引用
func tryFix1(t *testing.T) {
	for _, item := range ps {
		defer func() {
			i := item
			t.Log(i.Name)
		}()
	}
}

// 失败原因：对于 defer 后面的 func 来说，闭包的依旧是 item，而不是每一次的 item.Name
func tryFix2(t *testing.T) {
	for _, item := range ps {
		defer func() {
			name := item.Name
			t.Log(name)
		}()
	}
}

func fix1(t *testing.T) {
	for _, item := range ps {
		item := item
		defer func() {
			t.Log(item.Name)
		}()
	}
}

func fix2(t *testing.T) {
	for _, item := range ps {
		name := item.Name
		defer func() {
			t.Log(name)
		}()
	}
}

// 为什么能解决：不要被冗杂的语法迷惑了双眼，循环体只要看成一个函数的调用即可
// 方法传参是对变量值的引用，闭包是对变量的引用
func fix3(t *testing.T) {
	for _, item := range ps {
		defer func(name string) {
			t.Log(name)
		}(item.Name)
	}
}
