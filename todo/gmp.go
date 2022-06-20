package todo

/*
【GMP】
G（Goroutine）协程，虚拟概念
P（Processor）处理器，虚拟概念
M（Thread）内核线程

【G 的位置】
全局队列：存放等待运行的 G
P 的本地队列：存放等待运行的 G、不能超过 256 G、优先，如果本地队列放不下了，才放全局里

【P 的数量】
程序启动时创建 P，最多 GOMAXPROCS 个（可配置）
配置：环境变量 $GOMAXPROCS 或 在程序中通过 runtime.GOMAXPROCS(int) 设置

这里面有一个概念需要解释澄清，限制使用的 CPU 数和这里的 P，并不是两个概念，正是通过限制了 P 的数量，起到了限制 CPU 数的作用
就是说，无论你线程有多少个，都是要跑 P 本地队列中的 G 的，也就是 P 的数量就是同一时间能运行的任务数量的最大值

TODO
  这种角度的理解也不一定对，等了解的更加深刻后，在对这里的概念进行澄清解释；
  因为上面的概念也有可能说是，限制了使用的核数，自动实现了限制 P 的数量；
  runtime.GOMAXPROCS 上的源码注释也写到，是限制 CPU 数量的

【M 的数量】
当前操作系统分配到当前 Go 程序的内核线程数
Go 语言本身限定 M 的最大值是 10000
配置：retime/debug 中 SetMaxThreads

M 阻塞就创建，空闲就回收或者睡眠

【调度器设计策略】

[复用线程]
避免频繁的创建、销毁线程，而是对线程的复用
work stealing：当本地线程无可运行的 G 时，尝试从其他线程绑定的 P 窃取 G，而不是销毁线程
hand off：当本地线程因为 G 进行系统调用阻塞时，线程释放绑定的 P，把 P 转移给其他空闲的线程执行

[利用并行]
最多有 GOMAXPROCS 个线程分布在多个 CPU 上同时执行

[抢占]
coroutine 中，要等待一个协程主动让出 CPU，才执行下一个协程，在 Go 中，一个 Goroutine 最多占用 CPU 10ms、防止其他 Goroutine 被饿死

[全局 G 队列]
M 执行 work stealing 从其他 P 偷不到 G 时，它可以从全局 G 队列中获取 G
*/
