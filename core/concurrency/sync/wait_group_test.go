package sync

import (
	"sync"
	"testing"
	"time"
)

// 点1：描述 sync.WaitGroup{} 的语法 和 应用场景（类比 Java 中的 CountDownLatch 和 CyclicBarrier 去理解）
// 点2：循环中的闭包引用外部变量的问题：for range 初始化的变量只有一个，局部定义的函数引用的不是每次执行的函数时变量的值 - 没有进行拷贝，而是引用的该变量

// 闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)
func TestWaitGroup(t *testing.T) {
	// slice 的值类型换成基础的值类型也是一样的
	theMapSlice := []map[string]string{{"1": "1"}, {"2": "2"}, {"3": "3"}}

	group := sync.WaitGroup{}
	for _, theMap := range theMapSlice {
		group.Add(1)
		go func() {
			defer group.Done()
			// 意外：打印的都是 [3:3]，这个就是点题，点了第 2 点
			// 其实 goland 已经有提示了：Loop variable 'theMap' captured by func literal

			// 实际不加上下边这一句，效果也是一样的，因为 for 的速度 >> 启动一个 Goroutine 的速度
			// time.Sleep(time.Second)

			// 得到希望的结果可以把 time.Sleep 加在循环体里
			// 实际开发生产中，不可能采用这种做法，并发优势全无，毫无性能可言，解决方法：
			// 1、作为函数的参数（闭包不是，函数传参才是真正意义的值拷贝）
			// 2、每次循环都创建一个独立的新变量来规避这个问题
			println(theMap)
		}()

		time.Sleep(time.Second)
	}
	group.Wait()
	println("都执行完了")
}

// PS：局部函数的闭包不是对目标的一份拷贝么？因为之前权重指标树每一个都去初始化了，而这里表现的是却是公用一个
// 分析了一下，发现权重指标树那里是并发访问同一资源带来的问题；当初在那里没有意识到，还以为局部定义的函数会对引用的资源进行了拷贝，所以认为每一份都单独初始化了
// 就是并发读不会有任何问题，而并发写就肯定有问题 - 初始化就是写操作
