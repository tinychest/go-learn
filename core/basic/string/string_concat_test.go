package string

import (
	"fmt"
	"strings"
	"testing"
)

/*
【做字符串拼接的性能表现上】
 预分配内存(cap)的[]byte > []byte ~> strings.Builder ~> bytes.Buffer
 推荐使用预分配内存（strings.Builder.Grow(int)）的 strings.Builder 性能最好（比预分配内存的 []byte 少了一次 []byte 到 string 的转化，详见 Builder.String）

 不要使用 + 拼串
 不要使用 fmt.Sprintf 拼串
 因为拼串的本质是，不断申请拼串后大小的新的内存空间，来存储新的拼串结果

【strings.Builder】
底层就是一个字节数组，WriteString 就是直接调用的 append

[WriteString 的扩容]：切片的扩容机制
[Grow 的扩容]：n 小于 cap - len，底层切片就不扩容，否则扩容到原来 cap 的两倍 + n

【其他】
len 原生函数返回的是参数的 字节数
strings.Builder.Grow 的参数也是 字节数
编码差异，详见 basic_test.go

字符串拼接的性能比较，还可以参见：https://geektutu.com/post/hpg-string-concat.html
参见 performance/string_concat_test.go
*/

func TestGrow(t *testing.T) {
	// growTest1(t)
	// growTest2(t)
	// growBestPracticeTest(t)
}

func growTest1(t *testing.T) {
	b := strings.Builder{}

	// 初始，len 0 cap 0；扩容后 len 0 cap 2
	b.Grow(2)
	printBuilderInfo(&b)
	// 写入 2 字节数据后 len 2 cap 2
	b.WriteString("12")
	printBuilderInfo(&b)
	// 写入 8 字节数据后，len 10 cap 16
	b.WriteString("12345678")
	printBuilderInfo(&b)
}

// Grow 不是直接扩容，这个方法的调用含义是，保证 Builder 能写入指定字节数的数据，底层数据而不进行扩容
func growTest2(t *testing.T) {
	b := strings.Builder{}

	b.Grow(10)
	b.WriteString("12345")
	// 经过上面操作，len 5 cap 10
	printBuilderInfo(&b)

	b.Grow(5)
	// 还想扩容 5，并没有起到任何效果，因为底层数组空间完全够再存储 5 个字节的数据
	printBuilderInfo(&b)

	b.Grow(7)
	// 有 5 个字节的空间，希望再存储 7 个字节的数据，实际扩容后的空间是多少？增 2？不是，是原有总容量 10 * 2 + Grow 参数 = 10 * 2 + 7 = 27
	printBuilderInfo(&b)
}

// 在某些情况下，能够大幅提升性能的做法，参见：strings.Join
func growBestPracticeTest(t *testing.T) {
	p := []string{"abc", "d", "e", "kk", "2357"}
	b := strings.Builder{}

	l := 0
	for i := 0; i < len(p); i++ {
		l += len(p[i])
	}
	b.Grow(l)
	for i := 0; i < len(p); i++ {
		b.WriteString(p[i])
	}
	printBuilderInfo(&b)
}

func printBuilderInfo(b *strings.Builder) {
	fmt.Printf("Len：%d，Cap：%d\n", b.Len(), b.Cap())
}
