package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 参见：https://ifeve.com/go-concurrency-atomic/

// 下列方法操作过程中，计算机中的任何 CPU 都不会进行其他的针对此值的读或者写；这样的约束是受到底层硬件的支持的
// atomic.LoadInt64() // 取
// atomic.StoreInt64() // 存（存储过程中，计算机中的任何 CPU 都不会进行其他的针对此值的读或者写；这样的约束是受到底层硬件的支持的）
// atomic.AddInt64() // 自增
// atomic.SwapInt64() // 赋值
// atomic.CompareAndSwapInt64(addr, old, new) // CAS *addr == old →  *add = new，并返回 *addr == old（相关词条：·乐观锁、ABA 问题、CPU 空转）

func TestInt64(t *testing.T) {
	const n = 10000
	wg := sync.WaitGroup{}
	wg.Add(n)

	var sum int64

	for i := 0; i < n; i++ {
		go func() {
			sum++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(sum)
}

func TestAtomicInt64(t *testing.T) {
	const n = 10000
	wg := sync.WaitGroup{}
	wg.Add(n)

	var sum int64

	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&sum, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(sum)
}

// 上面的可以理解为悲观锁，而这里是乐观锁
func TestAtomicCASInt64(t *testing.T) {
	const n = 10000
	wg := sync.WaitGroup{}
	wg.Add(n)

	var sum int64

	for i := 0; i < n; i++ {
		go func() {
			for {
				// 获取值
				// （当前例子测不出来，但是，在 32 位计算架构的计算机上写入一个 64 位的整数，如果这个写操作未完成时，有一个读操作被并发执行了，就很有可能会读取到一个只被修改了一半的数据，所以这里应该使用并发安全的方式读取）
				v := atomic.LoadInt64(&sum)
				// 看下获取的值有没有被修改，没有修改就去修改（重复这个过程）
				if atomic.CompareAndSwapInt64(&sum, v, v+1) {
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(sum)
}
