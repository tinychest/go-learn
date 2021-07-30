package timeout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
timer 在这里的意思是定时器
*/
func TestTimer(t *testing.T) {
	// timerTest1()
	// timerTest2()
	timeoutTest()

	// time.NewTicker() *time.Ticket
	// time.NewTimer() *time.Timer
	// time.AfterFunc(d, f) *time.Timer
	// time.After(d) <-chan time.Time
}

/*
Reset 还需要搞得更透彻

如果明确 time 已经 expired，并且t.C已经被取空，那么可以直接使用 Reset；
如果程序之前没有从 t.C 中读取过值，这时需要首先调用 Stop()，如果返回 true，说明 timer 还没有 expire，stop 成功删除 timer，可直接 reset；
如果返回 false，说明 stop 前已经 expire，需要显式 drain channel
*/

// timerTest1 循环定时（for-timer-wait）
// 注意，你可以直接写一个简单 case <-time.After(interval) 但那会产生大量的垃圾
func timerTest1() {
	interval := time.Second
	timer := time.NewTimer(interval)
	defer timer.Stop()

	for {
		// 不要直接去思考这个 Stop 的作用，完全是因为 Reset 方法加的，详见 Reset 方法的源码注释
		if !timer.Stop() {
			// 尝试停止 timer，如果停止失败了，就主动重制 timer
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(interval)

		// timer 到点了，会向 C 中发送一个时间信号
		<-timer.C
		fmt.Println("one second later...")
	}
}

// timerTest2 循环定时（ticket）
func timerTest2() {
	interval := time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("one second later...")
	}
}

func timeoutTest() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// cancel()
	fmt.Println(ctx.Err())

	select {
	case <-ctx.Done():
		fmt.Println("one second later...")
	}

	fmt.Println(ctx.Err())
}
