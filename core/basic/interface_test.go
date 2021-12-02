package basic

import (
	"fmt"
	"testing"
)

// 接口定义
type i interface {
	hello() i
}

// 等价接口
type i2 interface {
	hello() i
}

// 标准实现
type s struct{}
// 实现1
type s1 struct{}
// 实现2 虽然调用表现上，就好像实现了 i 接口，实际上并不是
type s2 struct {
	hello func() i
}

func (s s) hello() i {
	return s
}

func (s s1) Hello() s1 {
	return s
}

func case1() {
	var s = new(s)
	var i1 i = s
	var i2 i2 = s

	// 相同方法定义的接口，可以互相赋值（当然，大的可以赋值给小的）
	i2 = i1
	i1 = i2
}

func case2(*testing.T) {
	var theP interface{} = new(s)

	// 显式接口
	if theI, ok := theP.(i); ok {
		theI.hello()
	} else {
		fmt.Println("not impl i interface")
	}

	// 匿名接口（也可以理解成是否实现了指定的方法签名）
	if theI, ok := theP.(interface{ Hello() }); ok {
		theI.Hello()
	} else {
		fmt.Println("not impl Hello Func")
	}
}