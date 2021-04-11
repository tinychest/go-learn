package context

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const playerName = "大胖"

// 吃鸡大胃王比赛
// 场景描述：选手直接开始吃 → 选手一轮一轮的吃，并将每一轮的结果告诉裁判，裁判接收到后才能开启新一轮 → 裁判决定比赛什么时候结束，选手会收到通知

// 通过 context.WithCancel(context.Background())
// 返回的 ctx 就是创建的上下文，子协程必须监听 ctx.Done()
// 返回的 cancel函数，调用时会触发 ctx.Done()
func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// 主协程：宣布开始吃鸡
	println("裁判：比赛开始...")
	eatNumChannel := eatChicken(ctx)

	// 主协程：计数，如果到了最大 10 轮，就终止比赛
	round := 0
	eatTotalCount := 0
	for eatCount := range eatNumChannel {
		round++
		eatTotalCount += eatCount
		fmt.Printf("裁判：第 %d 轮，%s 选手吃掉 %d 只炸鸡\n", round, playerName, eatCount)
		if round == 10 {
			println("裁判：比赛结束...")
			// 调用 cancel函数 就会触发 ctx.Done()
			cancel()
			break
		}
	}

	println("裁判：正在统计结果...")
	// 需要等待一下，不然子协程收到结束，都来不及处理
	time.Sleep(1 * time.Second)
	fmt.Printf("裁判：%s 选手一共吃了 %d 只炸鸡\n", playerName, eatTotalCount)
}

func eatChicken(ctx context.Context) <-chan int {
	counterChannel := make(chan int)
	eatTotalCount := 0
	round := 0

	go func() {
		for {
			// 新一轮，开始吃
			round++
			eatCount := rand.Intn(5) + 1
			eatTotalCount += eatCount

			select {
			// 裁判宣布比赛结束
			case <-ctx.Done():
				fmt.Printf("%s：哈？结束了？那我刚吃的 %d 还算不算了？\n", playerName, eatCount)
				return
			// 选手告诉裁判本轮吃了多少
			case counterChannel <- eatCount:
			}
		}
	}()

	return counterChannel
}
