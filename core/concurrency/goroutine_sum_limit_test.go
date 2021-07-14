package concurrency

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"
)

var maxGoroutineCurrencySum = runtime.GOMAXPROCS(8)

func TestGoroutineSumLimit(t *testing.T) {
	// theFuncForFunc := func(index int) func() {
	//     return func() {
	//         time.Sleep(3 * time.Second)
	//         sum := 0
	//         for i := 1; i <= index; i++ {
	//             sum += i
	//         }
	//         fmt.Printf("%d 的阶和为：%d\n", index, sum)
	//     }
	// }

	theFuncForFunc2 := func(index int) func() error {
		return func() error {
			time.Sleep(3 * time.Second)

			if index == 9 {
				return errors.New("我不干了再见")
			}

			return nil
		}
	}

	// LimitMaxRoutineRunTest(theFuncForFunc, 100)
	LimitMaxRoutineRunPrefTest(theFuncForFunc2, 100)
}

// 限制最大并行的 Goroutine 数 - 版本一
// 实现原理：很简单，就是利用通道的缓存
// theFuncForFunc：充当 Goroutine 获取要执行任务的任务池
// totalCount：任务池中总共的任务数
func LimitMaxRoutineRunTest(theFuncForFunc func(index int) func(), totalCount int) {
	maxConcurrentLimitChannel := make(chan bool, maxGoroutineCurrencySum)

	for index := 0; index < totalCount; index++ {
		select {
		case maxConcurrentLimitChannel <- true:
			fmt.Printf("第 %d 次运行...\n", index+1)
			go func() {
				theFuncForFunc(index)()
				<-maxConcurrentLimitChannel
			}()
		}
	}
}

// 限制最大并行的 Goroutine 数 - 版本二
// 任务执行失败的处理：1、停止继续发放任务执行（✔） 2、终止正在运行的 Goroutine（？）
func LimitMaxRoutineRunPrefTest(theFuncForFunc func(index int) func() error, totalCount int) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxConcurrentLimitChannel := make(chan bool, maxGoroutineCurrencySum)

	theFunc := func(callFunc func() error, index int) {
		err := callFunc()

		if err != nil {
			// 将错误信息存储到 context 中
			ctx = context.WithValue(ctx, "index", index)
			ctx = context.WithValue(ctx, "error", err)
			// 通知任务池的结束任务发放
			cancel()
		}

		// 任务完成
		<-maxConcurrentLimitChannel
	}

	for index := 0; index < totalCount; index++ {
		select {
		// 错误信息优先处理
		case <-ctx.Done():
			index := ctx.Value("index")
			err := ctx.Value("error")
			fmt.Printf("%d 号任务在执行时，发生意外，执行该任务的 Goroutine：%s\n", index, err)
			return

		default:
			select {
			case maxConcurrentLimitChannel <- true:
				fmt.Printf("第 %d 次运行...\n", index+1)
				go theFunc(theFuncForFunc(index), index)
			}
		}
	}
}
