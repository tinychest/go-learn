package sync

import (
	"sync"
	"testing"
	"time"
)

// 描述 sync.WaitGroup{} 的语法 和 应用场景（类比 Java 中的 CountDownLatch 和 CyclicBarrier 去理解）
func TestWaitGroup(t *testing.T) {
	slice := []string{"1", "2", "3"}

	wg := sync.WaitGroup{}
	for _, elem := range slice {
		// 闭包问题（详见闭包模块）
		elem := elem
		wg.Add(1)
		go func() {
			defer wg.Done()
			println(elem)
		}()

		time.Sleep(time.Second)
	}
	wg.Wait()
	println("执行完毕")
}
