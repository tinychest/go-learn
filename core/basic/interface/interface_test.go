package _interface

import (
	"testing"
)

// Go 的接口类型是完全以方法为核心的（有点像废话）
//
// Go 中没有像 Java 中的 implements 关键字，即没有强制的概念；但是你希望编译器帮你检查一下，起到强制的效果，那可以通过在全局接口类型变量的声明方式
// var _ IXxx = Xxx

func TestInterface(t *testing.T) {
	equalTest(t)
	assertTest(t)
}

// 等价接口定义
type I1 interface {
	hello() I1
}
type I2 interface {
	hello() I1
}

// 实现1
type S1 struct{}

// 实现2
type S2 struct{}

// 虽然调用表现上，就好像实现了 I1 接口，实际上并不是
type s3 struct {
	hello func() I1
}

func (s S1) hello() I1 {
	return s
}

// 这样并不算实现了 I2
func (s S2) Hello() S2 {
	return s
}

// 只要包含的方法签名相同，两个接口类型可以被认为是相等的
func equalTest(t *testing.T) {
	var s1 = new(S1)
	var i1 I1 = s1
	var i2 I2 = s1

	// 相同方法定义的接口，可以互相赋值（当然，大的可以赋值给小的）
	i2 = i1
	i1 = i2
}

// 接口类型断言 和 匿名接口类型断言
func assertTest(t *testing.T) {
	var theP interface{} = new(S1)

	// 显式接口
	if theI, ok := theP.(I1); ok {
		theI.hello()
	} else {
		t.Log("not impl I1 interface")
	}

	// 匿名接口（也可以理解成是否实现了指定的方法签名）
	if theI, ok := theP.(interface{ Hello() }); ok {
		theI.Hello()
	} else {
		t.Log("not impl Hello Func")
	}
}
