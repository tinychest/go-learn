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

// 优化5：简单的都处理掉了，接下来往更加深层的方向进行分析。
// 结果集 []byte → string 会进行一次内存分配，但是可以采用 unsafe 黑科技来避免这个
// 19 → 18
func foo5(n int) string {
	var buf bytes.Buffer

	x := strconv.Itoa(n)
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

func TestAlloc5(t *testing.T) {
	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo5(rand.Int())
	})))
}
