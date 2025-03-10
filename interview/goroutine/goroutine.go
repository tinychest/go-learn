package goroutine

import (
	"fmt"
	"runtime"
)

// 这块内容将从一下角度出发进行陈述：
// > 补充 GMP 场景样例时，可以把里边对于 goroutine 的描述文案添加到这里
// - 计算机操作系统基础，按照 进程、线城、协程、goroutine 演进递进的顺序去描述
// - goroutine 的使用原则（如果你不知道开启的这个 goroutine 什么结束，那就不应该开启它）
// - goroutine 的经典泄露场景（慢等待；等待：channel、WaitGroup、Mutex）
// - 如何定位 goroutine 泄露问题（runtime.NumGoroutine、pprof）
// - 面试官会怎么考：让你写一个泄露的例子、给你一段代码，让你看下有什么问题
//
// [起源]
// 进程是计算机操作系统中 “资源分配” 的最小单位、线程是计算机操作系统中 “CPU 调度” 的最小单位。
// 协程比线程更加轻量。线程的创建和销毁、调度需要陷入到内核，而协程可以认为完全依赖用户空间创建、销毁和调度。
// 很多语言都开始都开始尝试支持协程，最有名的就是 goroutine（go 开发人员必须这样说 [狗头]）
//
// [使用原则]
// 一个 goroutine 开启了，就需要知道什么时候结束或者满足什么条件下结束； 否则，就会造成泄露问题；
// goroutine 的结束，大体上是两个方向 超时控制 和 信号通知，这两个方向的实现离不开配合 select 去使用

// 获取当前 goroutine 的数量
func numGoroutine() {
	fmt.Println(runtime.NumGoroutine())
}
