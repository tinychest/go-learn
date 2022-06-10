package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 悲观锁
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

// 乐观锁
func TestAtomicCASInt64(t *testing.T) {
	const n = 10000

	wg := sync.WaitGroup{}
	wg.Add(n)

	var sum int64
	for i := 0; i < n; i++ {
		go func() {
			for {
				// 获取值
				// （自己的操作系统是 64 位的，所以测不出来，在 32 位计算架构的计算机上写入一个 64 位的整数，如果这个写操作未完成时，有一个读操作被并发执行了，
				// 就很有可能会读取到一个只被修改了一半的数据，所以这里应该使用并发安全的方式读取）
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
