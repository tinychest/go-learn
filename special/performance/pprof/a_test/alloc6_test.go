package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"unsafe"
)

// 优化6：回顾优化 2 中，预分配空间有效果但是效果不大，只能判断是缓存区大小设置的不对，仔细看 x 不是 byte 类型，而是 string 类型，想知道有多少个字节，只要调用 len 就行了
// 18 → 3
func foo6(n int) string {
	var buf bytes.Buffer

	x := strconv.Itoa(n)

	buf.Grow(100000 * len(x))
	for i := 0; i < 100000; i++ {
		buf.WriteString(x)
	}
	sum := sha256.Sum256(buf.Bytes())

	var b = make([]byte, 0, int(sum[0]))
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := "abcdefghijklmnopqrstuvwxyz"[x%26]
		b = append(b, c)
	}
	return *(*string)(unsafe.Pointer(&b))
}

func TestAlloc6(t *testing.T) {
	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo6(rand.Int())
	})))
}
