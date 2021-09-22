package concurrency

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"
	"time"
)

var max = runtime.GOMAXPROCS(8)

type task func() error

// 这只是相当简单的测试，Goroutine 并没有复用
func TestGoroutineLimit(t *testing.T) {
	// 定义任务池
	forFunc2 := func(i int) task {
		return func() error {
			time.Sleep(3 * time.Second)

			if i == 9 {
				return errors.New("我不干了再见")
			}

			return nil
		}
	}
	// 限制最大 Goroutine 数执行
	limitPrefTest(forFunc2, 100)
}

// 限制最大并行的 Goroutine 数 - 版本二
// 任务执行失败的处理：1、停止继续发放任务执行（✔） 2、终止正在运行的 Goroutine（？）
func limitPrefTest(forFunc func(int) task, totalCount int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	limitSig := make(chan struct{}, max)

	theFunc := func(callFunc task, index int) {
		err := callFunc()

		if err != nil {
			// 将错误信息存储到 context 中
			ctx = context.WithValue(ctx, "index", index)
			ctx = context.WithValue(ctx, "error", err)
			// 通知任务池的结束任务发放
			cancel()
		}

		// 任务完成
		<-limitSig
	}

	for i := 0; i < totalCount; i++ {
		select {
		// 错误信息优先处理
		case <-ctx.Done():
			index := ctx.Value("index")
			err := ctx.Value("error")
			fmt.Printf("%d 号任务在执行时，发生意外，执行该任务的 Goroutine：%s\n", index, err)
			return
		default:
			select {
			case limitSig <- struct{}{}:
				fmt.Printf("第 %d 次运行...\n", i+1)
				go theFunc(forFunc(i), i)
			}
		}
	}
}
