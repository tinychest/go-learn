package timeout

/*
核心接口 context.Context

实现类0：context.emptyCtx 没有超时时间、不能取消、不存储额外信息，用来作为 context 树的根节点
一般不直接使用，而是通过该结构体的 Background() TODO() 等方法来获取接口实现实例（详见源码）

实现类1：cancelCtx
继承自 context，同时也实现了 Context 接口
WithCancel()：获取一个可取消的 context - cancelCtx

实现类2：timerCtx
增加了 timeout 机制
WithDeadline()： 返回一个基于 parent 的可取消 context

实现类3：valueCtx
存储键值对的数据
WithValue()：用于向 context 添加键值对，注意添加方式是以链状的方法
*/

/*
返回截止时间 是否有截止时间
Deadline() (deadline time.Time, ok bool)

绑定当前 Context 的任务被取消时，返回一个关闭的 channel；如果 Context 不会被取消，将返回 nil
Done() <-chan struct{}

Done 方法返回的通道没有关闭了，返回 nil；如果返回的通道关闭了，返回的 err 会包含原因
Err() error

获取存储的数据
Value(key interface{}) interface{}
*/
func _() {
	// 引出：在主 Goroutine 中，如果我们希望对开启的子 Goroutine 进行一个超时的控制
	// 一般会采用一个通道，定义一段监听该通道，并在接收到通道数据后视主 Goroutine 认为程序需要终止，进行子 Goroutine 关闭的逻辑

	// 说法：
	// 1、web 编程中，一个请求对应多个 goroutine 之间的数据交互
	// 2、超时控制
	// 3、上下文控制

	// panic: cannot create context from nil parent
	// ctx, cancel := context.WithCancel(nil)

	// 调用 cancel 其实就是关闭维系内部的通道，而通道关闭了，是能够立即返回对应类型的零值的
}