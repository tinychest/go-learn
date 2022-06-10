package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 参见：https://ifeve.com/go-concurrency-atomic/
//
// 下列方法操作过程中，计算机中的任何 CPU 都不会进行其他的针对此值的读或者写；这样的约束是受到底层硬件的支持的
// atomic.LoadInt64()  // 取
// atomic.StoreInt64() // 存（存储过程中，计算机中的任何 CPU 都不会进行其他的针对此值的读或者写；这样的约束是受到底层硬件的支持的）
// atomic.AddInt64()   // 自增
// atomic.SwapInt64()  // 赋值
// atomic.CompareAndSwapInt64(addr, old, new) // CAS *addr == old →  *add = new，并返回 *addr == old（相关词条：·乐观锁、ABA 问题、CPU 空转）
//
// 可以很简单了解 int64 的赋值不是原子操作，但 int32 呢，atomic.Int32 相关方法做了什么
// 参见：https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651452934&idx=2&sn=51d3c21ff731837d50e9061ce6d9d2e6&scene=21#wechat_redirect

const N = 10000

// [可以通过 go test --race . 检测运行样例中得到数据竞争]
func TestAtomic(t *testing.T) {
	t.Log(NormalTest()) // < 10000
	t.Log(AtomicTest()) // = 10000
}

func NormalTest() int64 {
	wg := sync.WaitGroup{}
	wg.Add(N)

	var sum int64
	for i := 0; i < N; i++ {
		go func() {
			sum++
			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}

func AtomicTest() int64 {
	wg := sync.WaitGroup{}
	wg.Add(N)

	var sum int64
	for i := 0; i < N; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	return sum
}
