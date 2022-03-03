package context

import (
	"context"
	"testing"
	"time"
)

func TestContextTimeout(t *testing.T) {
	test1(t)
	// test2(t)
	// test3(t)
	//
	// testSimple(t)
}

func test1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// cancel()
	t.Log(ctx.Err())

	select {
	case <-ctx.Done():
		t.Log("one second later...")
	}

	t.Log(ctx.Err())
}

// 注意！！！这里能做到 case <-ctx.Done() 一触发，程序就停止
// 吃鸡大胃王比赛2
// 场景描述：比赛开始，选手就开始拼命吃 → 裁判决定比赛什么时候结束
func test2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 主 Goroutine：宣布开始吃鸡
	t.Log("裁判：比赛开始...")
	round := 0

	interval := 500 * time.Millisecond
	timer := time.NewTimer(interval)
	for {
		timer.Reset(interval)

		select {
		// 裁判宣布比赛结束
		case <-ctx.Done():
			t.Logf("裁判：Stop!\n")
			t.Logf("选手A：No! I can eat more\n")
			return
		// 每 0.5 秒吃一只
		case <-timer.C:
			round++
			t.Logf("选手A：第 %d 只\n", round)
		}
	}
}

func testSimple(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	t.Log("裁判：比赛开始...")
	count := 0

	interval := 500 * time.Millisecond
	timer := time.NewTimer(interval)
	for range timer.C {
		select {
		case <-ctx.Done():
			t.Logf("裁判：Stop!\n")
			t.Logf("选手A：No! I can eat more\n")
			return
		// 每 0.5 秒吃一只
		default:
			count++
			t.Logf("选手A：第 %d 只\n", count)
		}

		timer.Reset(interval)
	}
}

// 不是说调用 context.WithTimeout 返回的 cancel 函数，等待指定的时间后再触发 ctx.Done()
// 而是调用 WithTimeout 就开始计时了
func test3(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	intChannel := make(chan int)

	// 生产者
	go func() {
		count := 0
		for range time.NewTicker(1 * time.Second).C {
			select {
			case <-ctx.Done():
				t.Log("Time's up!")
				close(intChannel)
				return
			default:
				count++
				intChannel <- count
				t.Logf("produce：%d\n", count)
			}
		}
	}()

	// 消费者
	for i := range intChannel {
		t.Logf("consume：%d\n", i)
	}
}
