package todo

import (
	"fmt"
	"testing"
)

type Parent struct {
}

func (p Parent) GetName() string {
	return "parent"
}

func (p Parent) SayHello() {
	fmt.Println("I am", p.GetName())
}

type Son struct {
	Parent
}

func (s Son) GetName() string {
	return "son"
}

func TestComposite(t *testing.T) {
	// Parent 实现了 SayHello，Son 没有实现，所以通过 son.SayHello 调用的是 Parent，但是里边调用的 GetName 也是 Parent
	// Go 不会像其他面向对象高级编程语言（如 Java），能够输出实际类型实现的方法
	var son Son
	son.SayHello()
}
