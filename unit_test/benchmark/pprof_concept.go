package benchmark

/*
- 主题：性能分析

- 简介
benchmark(基准测试) 可以度量某个函数或方法的性能，也就是说，如果我们知道性能的瓶颈点在哪里，benchmark 是一个非常好的方式。
但是面对一个未知的程序，如何去分析这个程序的性能，并找到瓶颈点呢？

- 核心
编译到程序中的 runtime/pprof 包
性能剖析工具 go tool pprof
*/

/*一、CPU 性能分析
CPU profiling，是最常见的性能分析类型
启动 CPU 分析时，运行时(runtime) 将每隔 10ms 中断一次，记录此时正在运行的协程(goroutines) 的堆栈信息
程序运行结束后，可以分析记录的数据找到最热代码路径(hottest code paths)

hottest code paths：一个函数在性能分析数据中出现的次数越多，说明执行该函数的代码路径(code path)花费的时间占总运行时间的比重越大
*/

/*二、内存性能分析
Memory profiling 记录堆内存分配时的堆栈信息，忽略栈内存分配信息

内存性能分析启用时，默认每1000次采样1次，这个比例是可以调整的。因为内存性能分析是基于采样的，因此基于内存分析数据来判断程序所有的内存使用情况是很困难的
*/

/*三、阻塞性能分析（Go 语言特有，因为通道特有）
block profiling 阻塞性能分析用来记录一个协程等待一个共享资源花费的时间，在判断程序的并发瓶颈时会很有用

阻塞的场景包括：

    在没有缓冲区的信道上发送或接收数据
    从空的信道上接收数据，或发送数据到满的信道上
    尝试获得一个已经被其他协程锁住的排它锁

一般情况下，当所有的 CPU 和内存瓶颈解决后，才会考虑这一类分析

/*四、锁性能分析
mutex profiling 和 block profiling 类似，更专注于因为锁竞争导致的等待或延时

- 1、CPU Profiling 详解
在 main 函数开头加上：
    pprof.StartCPUProfile(os.Stdout)
    defer pprof.StopCPUProfile()

配合：go run main.go > cpu.pprof

很简单，但是，一般来说，不建议将结果直接输出到标准输出，因为如果程序本身有输出，则会相互干扰，直接记录到一个文件中是最好的方式
    f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

配合：go run main.go

得到的 cpu.pprof 可以通过工具查看（需要安装 Graphviz）
    go tool pprof -http=:9999 cpu.pprof
访问：localhost:9999

也可以直接通过交互式命令行查看
    go tool pprof cpu.pprof
交互：
(pprof) top
(pprof) top --cum
(pprof) help

- 2、Memory Profiling 详解
详见 memory_profile_test.go，得到的 mem.pprof，通过如下命令解析：
    go tool pprof -http=:9999 ....../mem.pprof

从这张图中，我们可以看到 concat 消耗了 524k 内存，randomString 仅消耗了 22k 内存。
理论上，concat 函数仅仅是将 randomString 生成的字符串拼接起来，消耗的内存应该和 randomString 一致，但怎么会产生 20 倍的差异呢？
这和 Go 语言字符串内存分配的方式有关系（和 Java 如出一辙）。字符串是不可变的，因为将两个字符串拼接时，相当于是产生新的字符串，
如果当前的空间不足以容纳新的字符串，则会申请更大的空间，将新字符串完全拷贝过去，这消耗了 2 倍的内存空间。在这 100 次拼接的过程中，会产生多次字符串拷贝，从而消耗大量的内存

使用 Strings.Builder 后，内存消耗降为 1/8

- benchmark 也是可以生成 pprof 文件的
    -cpuprofile=$FILE
    -memprofile=$FILE, -memprofilerate=N 调整记录速率为原来的 1/N。
    -blockprofile=$FILE

- 事实上 go tool pprof 支持多种输出格式，直接敲该命令得到提示
*/
