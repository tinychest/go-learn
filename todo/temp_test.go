package todo

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Goland Regex：找寻后边不是跟着 bb 或 cc 的 aa "aa(?!bb|cc)"
// aabb aacc aadd
// aae aat aappo

func TestConst(t *testing.T) {
	// 不支持
	// const arr = [2]int{}
}

func TestCodePos(t *testing.T) {
	// 如何输出让 Goland 控制台能够识别代码位置，且点击能跳转的日志（蓝字、下划线）
	// 没有搜索相关内容，那就断点看下实际的输出内容，这样应该就可以知道了
	// - Log 方法本身会打印出代码位置
	// t.Log("1")
	// - http 地址本是蓝色下划线可点击的
	// t.Log("http:localhost:8080/)

	// 结论：Goland 只要能索引到对应位置，就会标注成可点击的样式
	fmt.Println("temp_test.go:25")
}

/* interface、func call playground */
type Killer interface {
	Kill()
}

type Jack struct{}

func (j Jack) Kill() {
	fmt.Println("slash slash slash")
}

func TestMess(t *testing.T) {
	Jack{}.Kill()
	Jack.Kill(Jack{})
	interface{ Kill() }.Kill(Jack{})
	((Killer)(Jack{})).(interface{ Kill() }).Kill()
}

/* 业务中遇到的打印信息的模棱两可 */
func TestPrintPit(t *testing.T) {
	t.Log([]string{"a", "b"})   // [a b]
	t.Log([]string{`"a b"`})    // ["a b"]
	t.Log([]string{`"a`, `b"`}) // ["a b"]

	t.Log(fmt.Sprintf("%+v", []string{"1", "2", "3"}))
	t.Log(fmt.Sprintf("%#v", []string{"1", "2", "3"}))
}

/* 很能会忽视的，被 go 关键字迷糊的问题（关键点：参数确定） */
func TestGo(t *testing.T) {
	ch1 := make(chan int)
	go fmt.Println(<-ch1) // 确认参数值时会发生什么呢？应该通过闭包函数去达到目标效果
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

/* 了解 bytes.Buffer Truncate 方法的错用，因为 os 包下也有一个 Truncate 方法 */
func TestTruncate(t *testing.T) {
	b := bytes.Buffer{}

	b.WriteString("123")

	b.Truncate(2)
	t.Log(b.String())
}

/* runtime.Caller demo */
func TestSyscall(t *testing.T) {
	caller, file, line, ok := runtime.Caller(0)
	t.Log(caller)
	t.Log(file)
	t.Log(line)
	t.Log(ok)
}
