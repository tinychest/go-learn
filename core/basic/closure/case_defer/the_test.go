package case_defer

import (
	"fmt"
	"go-learn/core"
	"testing"
)

var ps = []core.Person{
	{Name: "小明"}, {Name: "小红"}, {Name: "小光"},
}

func TestDefer(*testing.T) {
	// mistake()
	// tryFix1()
	// tryFix2()
	// fix1()
	// fix2()
	fix3()
}

func mistake() {
	for _, item := range ps {
		defer func() {
			fmt.Println(item.Name)
		}()
	}
}

// 失败原因：并未改变函数内对 for 局部变量的闭包引用
func tryFix1() {
	for _, item := range ps {
		defer func() {
			i := item
			fmt.Println(i.Name)
		}()
	}
}

// 失败原因：对于 defer 后面的 func 来说，闭包的依旧是 item，而不是每一次的 item.Name
func tryFix2() {
	for _, item := range ps {
		defer func() {
			name := item.Name
			fmt.Println(name)
		}()
	}
}

func fix1() {
	for _, item := range ps {
		item := item
		defer func() {
			fmt.Println(item.Name)
		}()
	}
}

func fix2() {
	for _, item := range ps {
		name := item.Name
		defer func() {
			fmt.Println(name)
		}()
	}
}

// 为什么能解决：不要被冗杂的语法迷惑了双眼，循环体只要看成一个函数的调用即可
// 方法传参是对变量值的引用，闭包是对变量的引用
func fix3() {
	for _, item := range ps {
		defer func(name string) {
			fmt.Println(name)
		}(item.Name)
	}
}
