package performance

import (
	"sync"
	"testing"
)

var once sync.Once

// sync.Once 这是 Go 标准库中相当重要的一个，某个方法需要保证在整个程序生命周期执行一次，那么可以将这个方法作为 Once 的参数，通过它来调用，这样可以保证只执行一次
// 实际常见应用场景：单例模模式、懒加载、初始化配置、数据库连接等，作用和 init 挺相似的

// 其他：Once 结构体有两个字段，而之所以将 done unit32 放在第一个，是因为作为 hot path 访问，能够减少 cpu 指令，能够提升性能

// 为什么放在第一个字段就能够减少指令呢？
// 因为结构体第一个字段的地址和结构体的指针是相同的，如果是第一个字段，直接对结构体的指针解引用即可
// 如果是其他的字段，除了结构体指针外，还需要计算与第一个值的偏移(calculate offset)
// 在机器码中，偏移量是随指令传递的附加值，CPU 需要做一次偏移值与指针的加法运算，才能获取要访问的值的地址
// 因此，访问第一个字段的机器代码更紧凑，速度更快
func TestSyncOnce(t *testing.T) {
	// 一、
	// ok
	// once := sync.Once{}; once.Do(initFunc)
	// 不 ok
	// (sync.Once{}).Do(initFunc)
	// ok
	// new(sync.Once).Do(initFunc)

	// 注意每次调用都是一个新的 once 那就和没有 once 是一样的

	// 二、一个 Once 实例只能绑定一个 func
	do1()
	do2()

	// 底层原理（atomic 保证了读取指定地址数的原子性，之所以要保证是因为不同的 cpu 可能不是 4 个字节，4 个字节一读，这个和 cpu 架构有关，早期的 cpu 可能就是一个字节一个字节的读）
	// atomic.LoadUint32(&o.done)
	// atomic.StoreInt32(&o.done) // 源码中，这个操作是加了锁的，为什么要加锁，源码也有解释说明
}

func do1() {
	once.Do(func() {
		println("do1")
	})
}

func do2() {
	once.Do(func() {
		println("do2")
	})
}
