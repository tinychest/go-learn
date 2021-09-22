package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

// 既然 优化1 没用，就去了
// 优化2：strconv.Itoa 的结果预先存起来 - 100158 → 159（减少了 10W 次，看来大循环里的操作是非常有必要去优化的）
// 效果：运行时间直接减半
func foo2(n int) string {
	var buf bytes.Buffer

	x := strconv.Itoa(n)
	for i := 0; i < 100000; i++ {
		buf.WriteString(x)
	}
	sum := sha256.Sum256(buf.Bytes())

	var b []byte
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)[x]
		b = append(b, c)
	}
	return string(b)
}

func TestAlloc2(t *testing.T) {
	fmt.Println("Allocs:", int(testing.AllocsPerRun(100, func() {
		foo2(rand.Int())
	})))
}