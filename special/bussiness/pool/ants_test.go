package pool

import (
	"github.com/panjf2000/ants/v2"
	"testing"
	"time"
)

// GitHub: https://github.com/panjf2000/ants

// ants goroutine pool 中：【pool】【worker】 的设计、自旋锁的使用、Option 的自定义初始化策略的写法、panicHandler 的自定义
// 源码非常有参考意义，但是和实际业务中的思考有一些出入，所以没有借鉴到真正想借鉴的，但是业务中的设计确实也没什么刺挑了
func TestAntsPool(t *testing.T) {
	taskGetter := func(index int) func() {
		return func() {
			// panic("") ants 没有 panic 中断任务机制，就是一直执行完
			t.Log(index)
		}
	}

	pool, _ := ants.NewPool(100000, ants.WithPreAlloc(true))
	defer pool.Release()

	// 提交 100 个任务
	for index := 0; index < 100; index++ {
		if err := pool.Submit(taskGetter(index)); err != nil {
			t.Fatal(err)
			return
		}
	}
	time.Sleep(time.Second)
}
