package string

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

/*
strings.Builder
- 实现了 io.Writer 接口
- 非并发安全，可能导致数据丢失
- 不要拷贝
*/

// TODO 黑科技返回的 string，不能修改？
func TestModifyStringBuildResult(t *testing.T) {
	// 可以修改
	s := "0"
	[]byte(s)[0] = 1

	// 不可以修改，直接崩溃
	b := strings.Builder{}
	b.WriteString("0")
	s = b.String()
	[]byte(s)[0] = 1

	// 已经存在值的 strings.Builder 不能拷贝，调用 Grow、Write、WriteRune、WriteString 会发生 panic，但是 Reset、Len、String 是可以的
	b2 := b
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
