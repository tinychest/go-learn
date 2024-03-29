/*
Package pool

以恒定数量的 Goroutine 执行指定的批量任务（防止没有限制的创建 Goroutine 的低效做法）

n 发送者，1 接收者 的场景下，应该通过关闭额外的辅助通道来结束，以保证所有发送者都能正常退出

TODO 拓展
咋看对比测试中，特定场景，ants 的性能表现还不如自定义 pool，但是 ants 中有许多优秀的设计，毋庸置疑；因此，以后自定义 pool 可能会做出改进的点：

一、每次使用都需要 NewCenter，这样的实例利用率是很低的，GC 负担重，应该定义一个复用机制
- 已经关闭的通道不可能再打开，因此当前的设计，避免不了部分元素的重复创建（ants 的设计理念是 pool 是常驻的，可以随时向里边提交任务）

二、为了灵活、可拓展的管理 Worker 的数量，也应当将 Worker 也像任务一样进行 “池化管理”
（任务分发器获取到任务，然后去获取可用的 Worker，将任务丢到他的通道里）
- 可以维护一个实时获取活跃的 Worker 数量的方法
- Worker 是可以复用的（sync.Pool）

三、ants 为每个任务定义了超时机制；具体设置方式则是通过 opts 的高可拓展性的方式设置的（任务超时了，会将 Worker 置为超时状态）
自己目前是为一批任务定义的超时时间
*/
package pool
