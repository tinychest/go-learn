package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 吃鸡大胃王比赛2
// 场景描述：比赛开始，选手就开始拼命吃 → 裁判决定比赛什么时候结束
func TestTimeOutContext(t *testing.T) {
	ctx, timeoutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer timeoutCancel()

	// 主协程：宣布开始吃鸡
	println("裁判：比赛开始...")
	round := 0

	func() {
		for {
			select {
			// 裁判宣布比赛结束
			case <-ctx.Done():
				fmt.Printf("裁判：Stop!\n")
				fmt.Printf("%s：No! I can eat more\n", playerName)
				return
			// 每 0.5 秒吃一只
			case <-time.After(500 * time.Millisecond):
				round++
				fmt.Printf("%s：第 %d 只\n", playerName, round)
			}
		}
	}()

}

// 不是说调用 context.WithTimeout 返回的 cancel函数，等待指定的时间后再触发 ctx.Done()
// 而是调用 WithTimeout 就开始计时了
func TestTimeoutContext2(t *testing.T) {
	intChannel := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// 消费者：
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				println("child process interrupt...")
				cancel()
				close(intChannel) // 没有这个，main 的生产者会因为没有消费者消费，而被卡住
				return
			default:
				fmt.Printf("receive：%d\n", <-intChannel)
			}
		}
	}()

	// 生产者：10 个轮流生产的生产者
	for i := 0; i < 10; i++ {
		fmt.Printf("produce：%d\n", i)
		intChannel <- i
	}

	// defer close(intChannel)
	// defer cancel()
}
