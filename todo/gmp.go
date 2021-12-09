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

【M 的数量】
当前操作系统分配到当前 Go 程序的内核线程数
Go 语言本身限定 M 的最大值是 10000
配置：retime/debug 中 SetMaxThreads

M 阻塞就创建，空闲就回收或者睡眠
*/
