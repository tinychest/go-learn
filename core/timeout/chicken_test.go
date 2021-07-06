package timeout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const playerName = "胖虎"

// 注意！！！这里能做到 case <-ctx.Done() 一触发，程序就停止
// 吃鸡大胃王比赛2
// 场景描述：比赛开始，选手就开始拼命吃 → 裁判决定比赛什么时候结束
func TestTimeOutContext(t *testing.T) {
	ctx, timeoutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer timeoutCancel()

	// 主协程：宣布开始吃鸡
	println("裁判：比赛开始...")
	round := 0

	interval := 500 * time.Millisecond
	timer := time.NewTimer(interval)
	for {
		timer.Reset(interval)

		select {
		// 裁判宣布比赛结束
		case <-ctx.Done():
			fmt.Printf("裁判：Stop!\n")
			fmt.Printf("%s：No! I can eat more\n", playerName)
			return
		// 每 0.5 秒吃一只
		case <-timer.C:
			round++
			fmt.Printf("%s：第 %d 只\n", playerName, round)
		}
	}
}

// 不是说调用 context.WithTimeout 返回的 cancel函数，等待指定的时间后再触发 ctx.Done()
// 而是调用 WithTimeout 就开始计时了
func TestTimeoutContext(t *testing.T) {
	intChannel := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer close(intChannel)
	defer cancel()

	// 生产者
	go func() {
		index := 0
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				println("Time's up!")
				close(intChannel)
				return
			default:
				index++
				intChannel<-index
				fmt.Printf("produce：%d\n", index)
			}
		}

	}()

	// 消费者
	for i := range intChannel {
		fmt.Printf("consume：%d\n", i)
	}
}
