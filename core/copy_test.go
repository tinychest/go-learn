package core

import (
	"fmt"
	"testing"
)

type box struct {
	Name string
	Age  int
}

func (b box) Copy() *box {
	return &b
}

func TestCopy(t *testing.T) {
	b := box{
		Name: "盒子",
		Age:  10,
	}
	bCopy := b.Copy()
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", bCopy)

	// 指针类型也可以复制一个新的
	bPtr := &b
	bPtrCopy := bPtr.Copy()
	fmt.Printf("%p\n", bPtr)
	fmt.Printf("%p\n", bPtrCopy)
}
