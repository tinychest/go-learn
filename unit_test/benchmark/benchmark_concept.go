package benchmark

/*
主题：基准测试、性能测试

参照：https://geektutu.com/post/hpg-benchmark.html

- 简介
benchmark 翻译为 基准测试，Go 语言标准库内置的 testing 测试框架提供了基准测试的功能，从而让我们很轻易的进行性能测试

- 注意
性能测试受环境的影响很大，为了保证测试的可重复性，在进行性能测试时，尽可能地保持测试环境的稳定。

    1.机器处于闲置状态，测试时不要执行其他任务，也不要和其他人共享硬件资源。
    2.机器是否关闭了节能模式，一般笔记本会默认打开这个模式，测试时关闭。
    3.避免使用虚拟机和云主机进行测试，一般情况下，为了尽可能地提高资源的利用率，虚拟机和云主机 CPU 和内存一般会超分配，超分机器的性能表现会非常地不稳定。

- 要求
同单元测试，benchmark 样例的文件名需要以 _test 结尾

- go 测试命令介绍
go test <module name>/<package name>：运行某个 package 内的所有测试用例
go test .：运行当前 package 内的用例
go test ./<package name>：运行当前子 package 的所有用例
go test ./...：测试当前包下的所有 packag

go test 命令默认不运行 benchmark 用例，想要运行需要加上 -bench 参数
---《-bench 参数》---
支持传入一个正则表达式，匹配到的用例才会得到执行，例如只运行以 Fib 结尾的 benchmark 用例：
     go test -bench="Fib$" .

记一个大坑：现象贼奇怪，只要给 -bench 指定了参数，测试结果的控制台内容就贼少，最终发现是表示 -bench 参数内容需要用双引号

- benchmark 工作原理
b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行
b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快

- 实例，在当前包下执行 `go test -bench .` 得到：
goos: windows                                             - 操作系统
goarch: amd64                                             - cpu 架构
pkg: go-learn/unit_test/benchmark                         - 测试包
BenchmarkFib-8               198           6050792 ns/op

BenchmarkFib 就是测试方法名
-8 即 GOMAXPROCS（默认等于 CPU 核数，可以通过 -cpu 参数改变 GOMAXPROCS 的值），事实上 -cpu 参数支持列表 -cpu=2,4，会按照列表参数的个数作为条件重复测试。该调整该参数可以测试 cpu 的核数是否会影响方法性能
198           6050792 ns/op 代表一共测试了 198 次，平均每次花费 6050792 ns（0.006 s）
最下边的 1.851 代表总共花费的时间

PASS
ok      go-learn/unit_test/benchmark    1.851s

- 提升准确度
---《-benchtime 参数》---
指定测试时间，不包括编译、执行、销毁的时间（默认 1s），例：-benchtime=5s
    指定测试次数，例：-benchtime=30x

---《-count 参数》---
指定测试的轮数（默认 1轮）

- 内存分配情况
---《-benchmem 参数》---
可以度量内存分配的次数。内存分配次数也性能也是息息相关的，例如不合理的切片容量，将导致内存重新分配，带来不必要的开销

在当前目录执行 `go test -bench="Slice" .` 查看实例的结果：
goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
BenchmarkSliceCap-8                           40          28725752 ns/op
BenchmarkSliceAppropriateCap-8                58          22154267 ns/op
PASS
ok      go-learn/unit_test/benchmark    3.452s

很明显，存在性能差距，即使是百万级别，性能影响也没有想象中的那么大

跑题了，现在加上 -benchmem 参数：`go test -bench="Slice" -benchmem .`
goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
BenchmarkSliceCap-8                           40          28850880 ns/op        45188429 B/op         43 allocs/op
BenchmarkSliceAppropriateCap-8                56          24536105 ns/op         8003604 B/op          1 allocs/op
PASS
ok      go-learn/unit_test/benchmark    3.597s


45188429 和 8003604：空间差 6 倍
43 和 1：内存分配了 43 次

- 时间复杂度分析
就是说不进行算法分析，进行理论上的时间复杂度，你是可以通过修改方法中例如某些循环的次数，来查看，时间对应的增长情况
比如，循环次数 10倍，执行消耗的时间也是 10倍，那么可以说明程序逻辑的时间复杂度是一个 O(n) 线性的

- ResetTimer
假如在执行核心逻辑体之前需要进行一段耗时的准备，例如测试的数据准备，准备的耗时时间不应该影响基准测试，那么可以在准备逻辑执行完毕后，执行：
    b.ResetTimer()
- StartTimer、StopTimer
    在你想要忽略计时的语句前后加上 b.StopTimer() b.StartTimer()
*/
