package lock

import (
	"sync"
	"testing"
)

// 主要对比 sync.Mutex（互斥锁） 和 sync.RWMutex（读写锁） 的性能差异
//
// 测试场景：
// 1、读多写少
// 2、读少写多
// 3、读写一致
//
// go test -bench="Lock" .
// goos: windows
// goarch: amd64
// pkg: go-learn/unit_test/benchmark/lock
// BenchmarkLockReadMore-8                1        1026946300 ns/op
// BenchmarkLockReadMoreRW-8             10         104306310 ns/op
// BenchmarkLockWriteMore-8               1        1019038100 ns/op
// BenchmarkLockWriteMoreRW-8             2         915579250 ns/op
// BenchmarkLockEqual-8                   1        1019341800 ns/op
// BenchmarkLockEqualRW-8                 2         509086700 ns/op
// PASS
// ok      go-learn/unit_test/benchmark/lock       8.565s
//
// 一、单位读写操作的时间降为 1 微秒
// 读写比为 9:1 时，读写锁的性能约为互斥锁的 8 倍
// 读写比为 1:9 时，读写锁性能相当
// 读写比为 5:5 时，读写锁的性能约为互斥锁的 2 倍
//
// 二、单位读写操作的时间降为 0.1 微秒
// 读写锁的性能优势下降到 3 倍，这也是可以理解的，因加锁而阻塞的时间占比减小，互斥锁带来的损耗自然就减小了
//
// 更多关于锁的：https://colobu.com/2018/12/18/dive-into-sync-mutex/
//
// 互斥锁有两种状态：正常状态和饥饿状态。
//
// 在正常状态下，所有等待锁的 goroutine 按照FIFO顺序等待。
// 唤醒的 goroutine 不会直接拥有锁，而是会和新请求锁的 goroutine 竞争锁的拥有。
// 新请求锁的 goroutine 具有优势：它正在 CPU 上执行，而且可能有好几个，所以刚刚唤醒的 goroutine 有很大可能在锁竞争中失败。
// 在这种情况下，这个被唤醒的 goroutine 会加入到等待队列的前面。 如果一个等待的 goroutine 超过 1ms 没有获取锁，那么它将会把锁转变为饥饿模式。
//
// 在饥饿模式下，锁的所有权将从 unlock 的 goroutine 直接交给交给等待队列中的第一个。新来的 goroutine 将不会尝试去获得锁，即使锁看起来是 unlock 状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。
//
// 如果一个等待的 goroutine 获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个；(2)它等待的时候小于1ms。它会将锁的状态转换为正常状态。
//
// 正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。
func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkLockReadMore(b *testing.B)    { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkLockReadMoreRW(b *testing.B)  { benchmark(b, &RWLock{}, 9, 1) }
func BenchmarkLockWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkLockWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }
func BenchmarkLockEqual(b *testing.B)       { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkLockEqualRW(b *testing.B)     { benchmark(b, &RWLock{}, 5, 5) }
