package performance

import (
	"sync"
	"testing"
)

// sync.Pool：保存和复用临时对象，减少内存分配，降低 GC 压力

// Go 语言从 1.3 版本开始提供了对象重用的机制，即 sync.Pool
// sync.Pool 是可伸缩的，同时也是并发安全的，其大小仅受限于内存的大小
// sync.Pool 用于存储那些被分配了但是没有被使用，而未来可能会使用的值
// 这样就可以不用再次经过内存分配，可直接复用已有对象，减轻 GC 的压力，从而提升系统的性能

// sync.Pool 的大小是可伸缩的，高负载时会动态扩容，存放在池中的对象如果不活跃了会被自动清理

// 实际性能测试见
func TestSyncPool(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	pool := sync.Pool{New: func() interface{} {
		return &person{name: "xm", age: 10}
	}}

	// 初始，池为空，就会调用一次 New 方法来创建一个
	p := pool.Get().(*person)
	// 假如对象用完，应该放回池子中
	pool.Put(p)
}
