package performance

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Cond：经常用在多个 goroutine 等待，一个 goroutine 通知（事件发生）的场景

// 1.其实在大多数场景，sync.WaitGroup 就够用了
// 2.如果一个通知一个等待，也完全可以通过 互斥锁 或 channel 就搞定了

// 但是假如有这样的场景，多个 goroutine 等待一个 goroutine 读取数据，读取好后，其他 goroutine 才能获取需要的数据继续执行
// 1.可以通过 cas 自选操作
// 2.可以让这些 goroutine 阻塞在一个通道上，当接收完毕，关闭这个通道
// 无论采用哪种方案，都会带来额外的复杂度，此时就适合使用 Go 标准库 sync 的 Cond 了

// 总结：Go 提供的类库，原理和使用上都很简单；Cond 就感觉没有特别封装，一切从简，使用上有些露骨的感觉
func SyncCondConcept() {
	lock := new(sync.Mutex)

	// NewCond：创建需要关联一个锁，sync.Mutex 或 sync.RWMutex 都行
	cond := sync.NewCond(lock)
	// Broadcast：广播唤醒所有
	cond.Broadcast()
	// Signal：唤醒一个 Goroutine
	cond.Signal()
	// Wait：会释放锁，挂起调用者所在的 goroutine，当被 Broadcast 或 Signal 唤醒，将取消阻塞，并重新上锁
	cond.Wait()
}

func TestSyncCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
}

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(time.Second)

	c.L.Lock()
	done = true
	c.L.Unlock()

	fmt.Println(name, "wakes all")

	c.Broadcast()
}
