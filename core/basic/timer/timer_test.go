package timer

import (
	"fmt"
	"testing"
	"time"
)

// time.NewTimer() *time.Timer
// time.AfterFunc(d, f) *time.Timer
// time.After(d) <-chan time.Time
// time.NewTicker() *time.Ticket

/*
timer 特性概括：
- timer 核心在于 C，这是一个只读的 time.Time 通道
- 代码中可以通过 for range 或者 单纯的同步从通道读取，来达到定时、延时的效果
- 不像 ticker 可以循环定时，timer 只能定时一次，但是可以通过手动 Reset 来达到重新计时的效果
- 调用 Stop，可以使未达到定时的 timer，不会再满足条件，注意 Stop 不会关闭 C（为了保证代码执行的正确性，详见源码）
- 通过手动控制 Reset，可以实现定时时间 = MAX(逻辑执行时间, 定时时间) 或者 逻辑执行时间 + 定时时间 两种效果
	（ticker 只能实现前者的效果）

timer.Stop：
如果明确 time 已经 expired，并且 t.C 已经被取空，那么可以直接使用 Reset；
如果程序之前没有从 t.C 中读取过值，这时需要首先调用 Stop()，如果返回 true，说明 timer 还没有 expire，stop 成功删除 timer，可直接 reset；
如果返回 false，说明 stop 前已经 expire，需要显式 drain channel
	// Stop does not close the channel, to prevent a read from the channel succeeding incorrectly.
	// To ensure the channel is empty after a call to Stop, check the return value and drain the channel.
	if !timer.Stop() {
		// timer has already expired or been stopped.
		<-timer.C
	}
	timer.Reset(interval)

	// 参照：https://studygolang.com/articles/9289
	// 如果 timer 可能存在没有被取 C 中值的时候，使用 timer.Reset 是需要注意的，应该按照源码中的提示进行 Reset 的调用；否则，直接调用即可
	// 源码注释确实没有写清楚，实际很多场景，如果按照源码提示来，反而造成了死锁，所以有人提出了 issue
	if !timer.Stop() {
		// timer has already expired or been stopped.
		select { case <-timer.C: default: }
	}
	timer.Reset(interval)
*/

const interval = time.Second

func TestTimer(t *testing.T) {
	// timerTest()
	timerStopTest()
	// tickerTest()
}

// 你可以直接写一个简单 case <-time.After(interval) 但那会产生大量的垃圾
func timerTest() {
	before := time.Now()

	timer := time.NewTimer(interval)
	for range timer.C {
		timer.Reset(interval)

		now := time.Now()
		fmt.Println(now.Sub(before))
		before = now
	}
}

func timerStopTest() {
	timer := time.NewTimer(interval)

	fmt.Println(1)
	time.Sleep(interval)
	fmt.Println(2)

	// 这里的 Reset 没有起到效果（这里就应该按照源码提示那样处理）
	if !timer.Stop() {
		<-timer.C
	}
	timer.Reset(interval)
	<-timer.C
	fmt.Println("3")

	timer.Reset(interval)
	<-timer.C
	fmt.Println("4")
}

// ticker.C 同样，返回的是一个只读通道（无法关闭），调用 ticker.Stop 不意味着关闭 ticker.C
func tickerTest() {
	// NewTicker 只会执行一次
	for range time.NewTicker(interval).C {
		fmt.Println("one second later...")
	}
}

// 定时器以 3 秒 ←→ 1 秒 交替形式的执行
func tickerTest2() {
	sum := 0
	before := time.Now()
	ticker := time.NewTicker(time.Second)

	for now := range ticker.C {
		fmt.Println(now.Sub(before))
		before = now

		sum++

		if sum%2 == 0 {
			ticker.Reset(time.Second)
		} else {
			ticker.Reset(3 * time.Second)
		}
	}
}
