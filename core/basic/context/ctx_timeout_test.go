package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const playerName = "胖虎"

func TestContextTimeout(t *testing.T) {
	test()
	test1()
	test2()
}

func test() {
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

// 注意！！！这里能做到 case <-ctx.Done() 一触发，程序就停止
// 吃鸡大胃王比赛2
// 场景描述：比赛开始，选手就开始拼命吃 → 裁判决定比赛什么时候结束
func test1() {
	ctx, timeoutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer timeoutCancel()

	// 主 Goroutine：宣布开始吃鸡
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

func test1Simple() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	println("裁判：比赛开始...")
	count := 0

	interval := 500 * time.Millisecond
	timer := time.NewTimer(interval)
	for range timer.C {
		select {
		case <-ctx.Done():
			fmt.Printf("裁判：Stop!\n")
			fmt.Printf("%s：No! I can eat more\n", playerName)
			return
		// 每 0.5 秒吃一只
		default:
			count++
			fmt.Printf("%s：第 %d 只\n", playerName, count)
		}

		timer.Reset(interval)
	}
}

// 不是说调用 context.WithTimeout 返回的 cancel 函数，等待指定的时间后再触发 ctx.Done()
// 而是调用 WithTimeout 就开始计时了
func test2() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	intChannel := make(chan int)

	// 生产者
	go func() {
		count := 0
		for range time.NewTicker(1 * time.Second).C {
			select {
			case <-ctx.Done():
				println("Time's up!")
				close(intChannel)
				return
			default:
				count++
				intChannel <- count
				fmt.Printf("produce：%d\n", count)
			}
		}
	}()

	// 消费者
	for i := range intChannel {
		fmt.Printf("consume：%d\n", i)
	}
}
