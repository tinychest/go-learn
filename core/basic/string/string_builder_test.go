package string

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

// strings.Builder
// - 实现了 io.Writer 接口
// - 非并发安全，可能导致数据丢失
// - 不要拷贝

func TestModifyStringBuildResult(t *testing.T) {
	s := "0"

	// 编译不通过
	// s[0] = 1

	// 转换切片可以修改（string 转成 []byte 实际上是创建了一个新的 string）
	[]byte(s)[0] = 1

	// 地址引用的直接转换的修改，将引发 go/src/runtime/signal_windows.go:260 - throw - fatal（比 panic 更严重）
	bs := *(*[]byte)(unsafe.Pointer(&s))
	defer func() {
		fmt.Println("尝试补救...")
		if e := recover(); e != nil {
			fmt.Println("补救成功")
		}
	}()
	bs[0] = 1

	// 已经存在值的 strings.Builder 不能拷贝，调用 Grow、Write、WriteRune、WriteString 会发生 panic，但是 Reset、Len、String 是可以的
	// （一些 sync 包下的实例也不希望被复制）
	b1 := strings.Builder{}
	b1.WriteString("1")
	b2 := b1
	b2.WriteString("2") // panic: strings: illegal use of non-zero Builder copied by value
}

func TestConcurrent(t *testing.T) {
	var b strings.Builder
	n := 0
	var wait sync.WaitGroup
	for n < 1000 {
		wait.Add(1)
		go func() {
			b.WriteString("1")
			n++
			wait.Done()
		}()
	}
	wait.Wait()
	fmt.Println(len(b.String()))
}
