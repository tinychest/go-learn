package main

import (
	"math/rand"
	"testing"
)

// 会得到 100158，这是一个相当大的数字了，接下来会通过各种办法修改 foo 方法的实现，来将内存分配次数降为 0
func TestAlloc(t *testing.T) {
	t.Log("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo(rand.Int())
	})))
}
