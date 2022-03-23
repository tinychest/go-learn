package pool

import (
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// GitHub: https://github.com/panjf2000/ants
//
// ants goroutine pool 中：【pool】【worker】 的设计、自旋锁的使用、Option 的自定义初始化策略的写法、panicHandler 的自定义
// 源码非常有参考意义，但是和实际业务中的思考有一些出入，所以没有借鉴到真正想借鉴的，但是业务中的设计确实也没什么刺挑了
func TestAntsPoolSimple(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		i := i
		err := ants.Submit(func() {
			wg.Add(1)
			defer wg.Done()

			t.Log(i)
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	wg.Wait()

	t.Logf("running goroutines: %d\n", ants.Running())
	t.Logf("finish all tasks.\n")
}

func TestTaskPanic(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	err := ants.Submit(func() {
		defer wg.Done()

		panic("gg")
	})
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()

	t.Log("still run")
	time.Sleep(time.Second) // 必须加等待，不然 ants 来不及处理 panic
}

// 既然，ants 只负责高效的执行任务，那任务错误，异常情况上，只能由开发者自己考虑
//
// - 建议一 假如任务量很大，我们可以先单独跑一个任务看下结果是否成功，再并发执行其他的任务；
// 因为一般一个任务执行发生了错误，其他的任务也肯定会发生错误；没有发生错误亦然
// 这是一种比较理想的最大性能考虑
//
// - 建议二 实际还是应该按照执行是否发生错误来进行考虑，即，在 defer WaitGroup.Done 时，也进行一下 recover 操作
func TestBestPractice(t *testing.T) {
	// 手动让下面提交任务失败
	ants.Release()

	var success uint32 = 1

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i := i

		wg.Add(1)
		err := ants.Submit(func() {
			defer func() {
				if err := recover(); err != nil {
					atomic.StoreUint32(&success, 0)
				}
				wg.Done()
			}()

			// 实际任务，假如发生了 panic
			panic(i)

			// TODO 假如发生了错误
			// 	atomic.StoreUint32(&success, 0)
		})
		if err != nil {
			// TODO 返回错误，结束当前方法函数，按照道理说 panic 都行
			t.Fatal(err)
		}
	}
	wg.Wait()

	// TODO 看下任务是否都执行成功（并发任务都执行完了，所以直接获取值，应该没问题）
	if success == 0 {
		t.Fatal("实际应用场景中，应当返回错误")
	}

	// 处理并发任务的结果
}
