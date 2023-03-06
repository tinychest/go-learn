package concurrency

/*
【通道】
[参考]: https://gfw.go101.org/article/channel.html
Goroutine 和 Channel 是 Go 的两个招牌特性，使得使用 Go 进行并发编程变得十分便利和有趣。

Go 语言设计团队的首任负责人 Rob Pike 对并发编程的一个建议是不要让计算通过共享内存来通讯，而应该让它们通过通讯来共享内存。
通道机制就是这种哲学的一个设计结果。

[数据结构] runtime/chan.go 33
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue（下面的数组是作为环形缓冲区使用的）
	buf      unsafe.Pointer // points to an array of dataqsiz elements（就是一个数组）
	elemsize uint16
	closed   uint32
	elemtype *_type // element type（元素类型的类型元数据）
	sendx    uint   // send index（环形起始指针）
	recvx    uint   // receive index（环形结束指针）
	recvq    waitq  // list of recv waiters（等待读取值的 Goroutine）
	sendq    waitq  // list of send waiters（等待发送值的 Goroutine）

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

[类型语法]
chan xxx：  定义既能接收也能发送的通道类型
<-chan xxx：定义只能接收的通道类型
chan<- xxx：定义只能发送的通道类型
chan 可以转成 <-chan、chan<- 类型，但是反之不行

len(chan)：    通道中的元素个数；一个没有缓存区的通道 len=cap=0
cap(chan)：    通道的缓冲区大小

- 无法判断一个只写通道是否关闭！判断可读通道是否关闭也不是任意场合下都通用的
- 不能关闭只读的通道（编译不通过）；一定是关闭一个可写的通道（cannot close receive-only channel）

- 已经关闭的通道：关闭（❌ panic）、发送数据（❌ panic）、读取数据（✔ 得到类型零值）
- nil 通道：    关闭（❌ panic）、发送数据（❌ 阻塞住）、读取数据（❌ 阻塞住）
- 可以从一个已经关闭的带有缓冲区的通道正常读取缓冲区里的数据

- for-range channel 语法，会在 channel 关闭且缓存区内容被读取完时结束

[通道传值类型考量]
在一个值被从一个协程传递到另一个协程的过程中，此值将被复制至少一次。 如果此传递值曾经在某个通道的缓冲队列中停留过，则它在此传递过程中将被复制两次。
一般说来，为了在数据传递过程中避免过大的复制成本，我们不应该使用尺寸很大的通道元素类型。 如果欲传送的值的尺寸较大，应该改用指针类型做为通道的元素类型。

- 有时，一个请求可能并不保证返回一份有效的数据。
对于这种情形，我们可以使用一个形如 struct{v T; err error} 的结构体类型或者一个空接口类型做为通道的元素类型以用来区分回应的值是否有效。

- 在一些通知场景中，即，我们并不关心通道回应的值，只关心是否发生。
这时，我们常常使用 struct{} 来作为通道的元素类型，因为它的尺寸为 0，能够给节约一些内存（不多）。
（其他，如定时通知，time.Sleep And send to channel? no, use the time.After）

- 一对一通知：一个无缓存通道；通知方完成后向通道塞值，被通知方监听通道
- 一对多通知：一个无缓存通道；通道方完成后关闭通道，被通知方监听通道（一对一也可以使用这种方式，该方式是实践中使用的最多的）
- 多对一通知：一般多使用 sync.WaitGroup

[关闭通道 - 详见通道关闭准则]
- 逻辑上会重复关闭通知通告，但添加对应 defer recover 操作 的暴力关闭
- 对通道的关闭进行相关的封装，上锁，标识位判断
*/
