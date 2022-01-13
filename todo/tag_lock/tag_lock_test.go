package tag_lock

import (
	"sync"
	"testing"
	"time"
)

// TODO 思考如何设计并发的测试样例来验证标识锁是安全的

// 标识锁机制的的标识定义（关于该类型的定义和 Go 类型的可比性紧密相关：https://mp.weixin.qq.com/s/_AYOAtNhPGZy4ttfsDDw8w）
type key struct {
	VerNo int
}

// 相同锁能锁住
func TestTagLock1(t *testing.T) {
	var verLock sync.Map

	k := key{VerNo: 2020}

	go func() {
		lock, _ := verLock.LoadOrStore(k, &sync.Mutex{})
		lock.(sync.Locker).Lock()
		t.Log("Goroutine 没有被锁住")
	}()

	// 休息一会，不然每次都是这里 main 执行完就结束了
	time.Sleep(time.Second)
	lock, _ := verLock.LoadOrStore(k, &sync.Mutex{})
	lock.(sync.Locker).Lock()
	t.Log("主 没有被锁住")
}

// 不同锁互不影响
func TestTagLock2(t *testing.T) {
	var verLock sync.Map

	go func() {
		lock, _ := verLock.LoadOrStore(key{2020}, &sync.Mutex{})
		lock.(sync.Locker).Lock()
		t.Log("Goroutine 没有被锁住")
	}()

	time.Sleep(time.Second)
	lock, _ := verLock.LoadOrStore(key{2021}, &sync.Mutex{})
	lock.(sync.Locker).Lock()
	t.Log("主 没有被锁住")
}
