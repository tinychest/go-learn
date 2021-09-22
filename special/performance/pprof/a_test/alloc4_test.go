package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// 优化4：接着向下看，b 很明显，进行预分配试一下
// 23 → 19
func foo4(n int) string {
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
	return string(b)
}

func TestAlloc4(t *testing.T) {
	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo4(rand.Int())
	})))
}
