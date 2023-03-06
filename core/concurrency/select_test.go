package concurrency

import (
	"testing"
)

// Select case must be receive, send or assign receive
//
// select 最重要的含义是能够进行 channel 的无阻塞式发送和读取，常和无条件 for 一起使用
// select {
//	case v := <-c1:
//		// do something
//  case c2 <- v:
//      // do something
//	default:
//      // do something
// }
//
// go 中还提供了能够实现定时的 timer，详见 timer_test.go

func TestSelect(t *testing.T) {
	// basicTest(t)
	doubleTest(t)
}

// 非常好的，能够说明 select 基本原理和用法的例子
// - select 会阻塞住，直至任意一个 case 满足条件，如果都不满足走的就是 default（如果定义了）
// - case 的选择是随机的
func basicTest(t *testing.T) {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go func() { c1 <- 1 }()
	go func() { c2 <- 2 }()
	go func() { c3 <- 3 }()

	// 如果没有上面的 Goroutine，那么下边每次执行的就是 default
	select {
	case v := <-c1:
		t.Logf("从通道1读取值：%d\n", v)
	case v := <-c2:
		t.Logf("从通道2读取值：%d\n", v)
	case v := <-c3:
		t.Logf("从通道3读取值：%d\n", v)
	default:
		t.Log("No match and default run")
	}
}

// go 支持直接将一个通道对接到另外一个通道，这个操作不要放在 select case 上
// 下面的案例你可能会认为 c1 和 c2 都至少会输出一次，实际：
// - c1 读取的值仅作为 c 的评估后丢弃
// - 评估第一个 case c 的可读性，进而读取 c1，导致死锁
func doubleTest(t *testing.T) {
	c := make(chan struct{}, 1)
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)

	c1 <- struct{}{}
	c2 <- struct{}{}

	for {
		select {
		case c <- <-c1:
			t.Log("c1")
		case <-c2:
			t.Log("c2")
		default:
			return
		}
	}
}
