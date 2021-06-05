package sync

import (
	"sync"
	"testing"
)

/*
一、说明

1.RWMutex 可以创建为其他结构体的字段；零值为解锁状态
2.RWMutex类型的锁也和协程无关，可以由不同的协程加读取锁/写入和解读取锁/写入锁
3.其实 sync.RWMutex 是从 sync.Mutex（同一时间只能有一个读或写的全局锁）提高多读少写场景性能的衍生

二、api

读锁 - 想要锁住需要等待写锁释放
instance.RLock()
instance.RUnlock()

写锁（全局锁）- 想要锁住需要等待所有读写锁释放
instance.Lock()
instance.Unlock()

三、注意

1.解锁不能再加锁之前，读锁的解锁次数不能大于读锁的加锁次数（待验证）
2.写锁的优先级大于读锁，即某一时刻，读锁和写锁都符合加索条件时，优先执行加写锁的逻辑
3.Unlock 这个不是通用的解锁方法，这是解全局锁，不能解读锁
*/

func TestRWMutex(t *testing.T) {

	rwMutexTest()
}

// 目前项目中的应用场景是这样的，因为一个 list 的 len、cap、data 的赋值创建不是原子操作，所以存在并发安全问题，所以通过如下的操作来保证 instance 中的 List 是赋值完毕的
// 下边的代码相当于写操作，演示了加写锁，所有的读操作都应该加上读锁，这里没有演示出来
func rwMutexTest() []string {
	type HeavyData struct {
		sync.RWMutex
		// 其他的重型数据
		List []string
	}

	instance := HeavyData{}

	// 对 list 进行初始化操作...
	// instance.initList()

	instance.Lock()
	list := instance.List
	instance.Unlock()

	return list
}
