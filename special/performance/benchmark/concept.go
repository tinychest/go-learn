package benchmark

/*
【benchmark（基准测试）】

[基准测试方法定义]
文件名 _test.go 结尾
方法名 Benchmark 打头
参数 b *testing.B
b.N 表示用例需要运行的次数，对于每个用例都是不同的，用例执行时间越短，这个数值会越大

【命令】
[语法]
go test  -bench [args...] <module name | package name（常用 .）>

示例：
cd xxx → go test -bench="Xxx$" -run=none -benchmem .

[参数]
-bench='<regex>'（默认不跑 benchmark，只有加了 -bench 参数才会运行 benchmark 用例，例如 'xxx$'、'^xxx'）
-cpu=2,4,...（使用多少 CPU 核数去执行用例，列表有多少个元素，结果就会有多少个用例运行结果）
-benchtime=<n>s（默认 1 秒，执行时间变长了，用例的执行次数就也相应变长）
-benchtime=<n>x（指定具体的执行次数）
-count=<n>（执行的轮数，有几轮就有几个结果）
-benchmem（结果中展示内存分配情况、分配次数）
-v（显示详细信息，将 Log、Logf 方法的结果也显示出来）

[结果]
Benchmark_Xxx-<使用核数> <执行次数> <每次执行花费的平均时间> <每次执行平均花费的内存> <每次执行的内存分配次数>

【其他】
- 测试不同的输入：应该抽象出传参测试方法，然后定义多个实例，以不同参数调用这个方法
- 测试应该将耗时的准备工作排除在外：在执行完耗时操作后，调用 b.ResetTimer()
  其他还有如：b.StopTimer()、b.StartTimer()
- go test 指令，会默认执行 *_test.go 中的所有测试方法（不限于 Benchmark，还有普通 test），
  可以通过 “-run=none” 参数来避免执行
*/
