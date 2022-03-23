package pool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// capacity：池容量 表示 ants 最多能创建的 goroutine 数量。如果为负数，表示容量无限制
// running：已经创建的 worker goroutine 的数量
//
// 需要同步等一批提交的任务执行完毕，需要开发者在提交任务的上下文中维护一个 sync.WaitGroup
// 需要注意 wg.Add 一定要放在 ants.Submit 的函数之外，否则可能会出现一个任务都没执行，WaitGroup.Wait 直接跳过去，导致后边的流程直接处理没有处理好的结果，造成意想不到的结果
//
// 注意样例的一些明显有问题的点
// - 设计上，即使是某个任务发生了 panic，ants 也会依旧执行完所有的任务
// - 样例中，关了默认的池子，假如程序代码中需要常驻一个 pool 用于提交任务，第一次用完关闭了，第二次提交任务将导致返回错误，这点需要注意
// - 样例中，定义的任务方法中应该通过 defer 去调用 WaitGroup.Done，否则假如方法发生 panic 了将导致内存泄漏（不会被直接判定成死锁）！详见下面的 TestMemoryLeak
// - 样例中，居然忽视提交任务返回的错误，假如 ants 发生了错误，将可能导致死锁！详见下面的 TestDeadlock（应当在发生了错误结束程序，也就是不能执行到 WaitGroup.Wait 方法）
//   在特定环境下，死锁可能就会变成内存泄漏！详见 TestDeadToLeak

/* 样例 1 */
func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

// 这是官方样例的写法（不完全一致）
func TestCommonPool(t *testing.T) {
	defer ants.Release()

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		_ = ants.Submit(func() {
			demoFunc()
			wg.Done()
		})
	}
	wg.Wait()

	t.Logf("running goroutines: %d\n", ants.Running())
	t.Logf("finish all tasks.\n")
}

func TestMemoryLeak(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	_ = ants.Submit(func() {
		panic("gg")
		wg.Done()
	})
	wg.Wait()
}

func TestDeadlock(t *testing.T) {
	ants.Release()

	var wg sync.WaitGroup

	wg.Add(1)
	_ = ants.Submit(func() {
		wg.Done()
	})
	wg.Wait()
}

func TestDeadToLeak(t *testing.T) {
	go func() {
		ants.Release()

		var wg sync.WaitGroup

		wg.Add(1)
		_ = ants.Submit(func() {
			wg.Done()
		})
		wg.Wait()

		t.Log("执行完毕") // 执行不到
	}()
	// 主 goroutine 阻塞住
	err := http.ListenAndServe(":8080", &http.ServeMux{})
	if err != nil {
		t.Fatal(err)
	}
}

/* 样例 2 */
var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	runTimes := 1000
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(1))
	}
	wg.Wait()

	t.Logf("running goroutines: %d\n", p.Running())
	t.Logf("finish all tasks, result is %d\n", sum)
}
