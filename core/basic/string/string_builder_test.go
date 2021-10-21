package string

import (
	"fmt"
	"strings"
	"testing"
)

// 所以做字符串拼接的性能表现上：
//  预分配内存(cap)的[]byte > []byte ~> strings.Builder（推荐） ~> bytes.Buffer
// 1.不要使用 + 的拼串 或 fmt.Sprintf
// 2.推荐使用预分配内存（strings.Builder.Grow(int)）的 strings.Builder 性能最好（比预分配内存的 []byte 少了一次 []byte 到 string 的转化）

// strings.Builder
// 底层就是一个字节数组，WriteString 就是直接调用的 append

// WriteString 的扩容：切片的扩容机制
// Grow 的扩容：n 小于 cap - len，底层切片就不扩容，否则扩容到原来 cap 的两倍 + n

// 注意 bytes.Buffer 和 strings.Builder 的扩容机制是不同的

func TestGrow(t *testing.T) {
	b := strings.Builder{}

	b.Grow(2)
	printBuilderInfo(&b)
	b.WriteString("12")
	printBuilderInfo(&b)
	b.WriteString("12345678")
	printBuilderInfo(&b)
}

func printBuilderInfo(b *strings.Builder) {
	fmt.Printf("Len：%d，Cap：%d\n", b.Len(), b.Cap())
}