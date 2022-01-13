package core

import (
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
	t.Logf("%p\n", &b)
	t.Logf("%p\n", bCopy)

	// 指针类型也可以复制一个新的
	bPtr := &b
	bPtrCopy := bPtr.Copy()
	t.Logf("%p\n", bPtr)
	t.Logf("%p\n", bPtrCopy)
}
