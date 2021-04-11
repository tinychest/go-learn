package context

import (
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	// panic: cannot create context from nil parent
	// ctx, cancel := context.WithCancel(nil)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			println("副：睡觉中...")
			time.Sleep(5 * time.Second)
			println("副：醒了")

			select {
			case <-ctx.Done():
				time.Sleep(5 * time.Second)
				println("会打印就好了")
			}
		}
	}()

	println("主：喂，结束了")
	cancel()
	println("主：总算收到了，再见")
	// 实验结论：调用了 cancel，主线程并不会卡住，即不等待 <- ctx.Done() 的消息被接收
}
