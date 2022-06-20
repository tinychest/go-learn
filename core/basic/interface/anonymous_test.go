package _interface

import "testing"

// 匿名的接口实现（换句话说，只有方法实现）
// 还真是那句经典的话，任何程序代码、实现或架构中的问题，都可以通过加一层来解决，一层不够就两层

type IHello interface {
	Hello()
}

type HelloWrap struct {
	Func func()
}

func (w HelloWrap) Hello() {
	w.Func()
}

func NewHelloWrap(f func()) HelloWrap {
	return HelloWrap{Func: f}
}

func HelloFunc() {
	println("hello")
}

func TestAnonymous(t *testing.T) {
	var _ IHello = NewHelloWrap(HelloFunc)
}
