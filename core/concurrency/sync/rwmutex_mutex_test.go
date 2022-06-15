package sync

import (
	"sync"
	"testing"
)

// 1.RWMutex 可以创建为其他结构体的字段；零值为解锁状态
// 2.RWMutex类型的锁也和 Goroutine 无关，可以由不同的 Goroutine 加读取锁/写入和解读取锁/写入锁
// 3.其实 sync.RWMutex 是从 sync.Mutex（同一时间只能有一个读或写的全局锁）提高多读少写场景性能的衍生
//
// 读锁 - 想要锁住需要等待写锁释放
// instance.RLock()
// instance.RUnlock()
//
// 写锁（全局锁）- 想要锁住需要等待所有读写锁释放
// instance.Lock()
// instance.Unlock()

// 读锁之间不互斥、写锁之间互斥、读锁和写锁之间互斥
func TestRR(t *testing.T) {
	rw := sync.RWMutex{}
	rw.RLock()

	rw.RLock()
	t.Log("rlock and rlock")
}

func TestRW(t *testing.T) {
	rw := sync.RWMutex{}
	rw.RLock()

	rw.Lock()
	t.Log("rlock and wlock")
}

func TestWW(t *testing.T) {
	rw := sync.RWMutex{}
	rw.Lock()

	rw.Lock()
	t.Log("wlock and wlock")
}
