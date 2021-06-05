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
}
