package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 以下方法示例都是基于某一个 Goroutine 中需要开启另外一个 Goroutine 去执行任务的简单测试场景逻辑

/*
Go 中可以无条件杀死一个 Goroutine 么？

- 答案：不可以，goroutine 只能自己退出，而不能被其他 goroutine 强制关闭或杀死
Java 中不推荐强行杀死线程的做法（早期的 Thread.stop 设计出来，并在之后的版本被废弃）；Go 中更是如此，从诞生之初就没有预留无条件结束 Goroutine 的方法
之所以没有，一个是从其他语言的最佳实践经验来说，二个是如果 Go 想要支持，可能面临着很大的难题：
    杀死一个 goroutine 设计上会有很多挑战，当前所拥有的资源如何处理？堆栈如何处理？defer 语句需要执行么？
    如果允许 defer 语句执行，那么 defer 语句可能阻塞 goroutine 退出，这种情况下怎么办呢？

- 官方：

GitHub 上都已经讨论过这个问题 https://github.com/golang/go/issues/32610

goroutine 被设计为不可以从外部无条件地结束掉，只能通过 channel 来与它通信。
也就是说，每一个 goroutine 都需要承担自己退出的责任。
(A goroutine cannot be programmatically killed. It can only commit a cooperative suicide.)

- 实际：
1、注意不能在超时的控制逻辑将 Goroutine 执行任务返回结果的通道给 close 了，已经关闭的通道，进行读或写都会 panic
2、题外话，如果存在 某个 Goroutine 中的主体逻辑就是读取某个通道，并且通道关闭，Goroutine 就应该结束的逻辑，Goroutine 中采用 for range 是一种比较好的实现逻辑
3、从下边的例子可以得到，通过 【带有 1 个缓存】 或者 【通过 Goroutine 任务中的 select case default】 都是确定超时后处理 Goroutine 任务结果的合理做法，但是更推荐 2 者
4、如果某个任务被拆分成多段，在每段任务开始执行前，判断一下是否超时，是比较合理的

*/
func TestTimeoutSimulate(t *testing.T) {
	// _, cancel := context.WithCancel(context.Background())
	// // 不会报错，也不会阻塞
	// cancel()
	// cancel()
}

// 方式1 - 基于以下语法特性
// 1、time.Sleep
// 2、select 每次只有一个 case 会执行，以及没有符合条件的 case，将走 default 的特性
func timeoutTest() {
	var (
		runtimeSecond = 4
		timeoutSecond = 3
		runtime       = time.Duration(runtimeSecond) * time.Second
		timeout       = time.Duration(timeoutSecond) * time.Second

		taskChannel    = make(chan int)
		timeoutChannel = make(chan int)

		taskFunc = func() {
			time.Sleep(runtime)
			taskChannel <- 1
		}

		timeoutFunc = func() {
			time.Sleep(timeout)
			timeoutChannel <- 1
		}
	)

	// 任务执行
	go taskFunc()

	// 超时
	go timeoutFunc()

	// 接收任务执行结果
	select {
	case intValue := <-taskChannel:
		fmt.Printf("从通道接收：%d\n", intValue)
	case <-timeoutChannel:
		fmt.Printf("达到 %d 超时时间\n", timeoutSecond)
	}
}

// 方式2 - 改进方式1中，存在的任务 Goroutine 没有结束的资源泄露情况（可以通过 t.Log(runtime.NumGoroutine()) 查看）
// Goroutine 没能退出是因为通道没有接收值的了，所以被阻塞住，所以解决核心就是防止阻塞住
func timeoutTest2() {
	var (
		runtimeSecond = 4
		timeoutSecond = 3
		runtime       = time.Duration(runtimeSecond) * time.Second
		timeout       = time.Duration(timeoutSecond) * time.Second

		// 方式一 为通道设置缓存
		taskChannel    = make(chan int, 1)
		timeoutChannel = make(chan int)

		taskFunc = func() {
			time.Sleep(runtime)
			// 方式二 添加 select case default
			taskChannel <- 1
		}

		timeoutFunc = func() {
			time.Sleep(timeout)
			timeoutChannel <- 1
		}
	)

	go taskFunc()
	go timeoutFunc()

	select {
	case intValue := <-taskChannel:
		fmt.Printf("从通道接收：%d\n", intValue)
	case <-timeoutChannel:
		fmt.Printf("达到 %d 超时时间\n", timeoutSecond)
	}
}

// 方式3 - 推出 time 包下 api 的用法，以及对逻辑完成时，Goroutine 结果通知通道的处理
func timeoutTest3() {
	var (
		finish      = false
		rwLock      sync.RWMutex
		taskChannel = make(chan int)
	)

	// 模拟耗时的业务逻辑
	go func() {
		time.Sleep(4 * time.Second)

		rwLock.RLock()
		defer rwLock.RUnlock()
		if !finish {
			taskChannel <- 1
		}
	}()

loop:
	for {
		select {
		// 接收结果
		case result := <-taskChannel:
			fmt.Printf("从通道接收：%d\n", result)
		// 超时逻辑
		case <-time.After(3 * time.Second):
			println("达到最大超时时间，主 Goroutine 结束")
			break loop
		}
	}

	rwLock.Lock()
	finish = true
	rwLock.Unlock()
}

// 最终 上边都是小 demo，实战中需要考虑诸多问题：方法调用者决定调用方法的超时、调用方法执行发生异常的意外中止等等，引出 context
// 详见 context 的相关 test
