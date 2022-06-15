package sync

import (
	"sync"
	"testing"
)

// RWMutex 相较于 Mutex 在读多写少场景下，具有更好的性能
// 使用上，都知道在并发读写的场景下，为 读 加上读锁，为 写 加上写锁，来保证数据的并发安全
//
// 这里演示一下，在并发读写的场景下，只在写的一方加写锁（读的一方不加读锁），会带来什么问题
// （把下面样例的读锁去掉，打印结果中的结果，会出现数字不连续的情况）
// 其实答案是显而易见的，带来的问题就是，假如某次数据我还没有读取完，就被修改了，导致该次读取的数据有逻辑上的问题
//
// - 通过 --race 参数得到资源竞争提示是必然的

func TestRWMutex(t *testing.T) {
	type Config struct {
		a []int
	}
	cfg := &Config{}
	rw := sync.RWMutex{}

	// 写
	go func() {
		i := 0
		for {
			i++
			rw.Lock()
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			rw.Unlock()
		}
	}()

	// 读
	for n := 0; n < 100; n++ {
		// rw.RLock()
		t.Logf("%#v\n", cfg)
		// rw.RUnlock()
	}
}
