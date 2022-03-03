package context

/*
核心接口 context.Context

[接口方法]
返回：截止时间 是否有截止时间（通过 WithDeadline、WithTimeout 获取得 ctx 才有）
Deadline() (deadline time.Time, ok bool)

绑定当前 Context 的任务被取消时，返回一个关闭的 channel；如果 Context 不会被取消，将返回 nil
Done() <-chan struct{}

Done 没有关闭，返回 nil；关闭了，返回关闭的原因：context was canceled | DeadlineExceeded
Err() error

获取存储的数据
Value(key interface{}) interface{}

[实现类]
实现类0：context.emptyCtx 没有超时时间、不能取消、不存储额外信息，用来作为 context 树的根节点
一般不直接使用，而是通过该结构体的 Background()

实现类1：cancelCtx
继承自 context，同时也实现了 Context 接口
WithCancel()：获取一个可取消的 context - cancelCtx

实现类2：valueCtx
存储键值对的数据
WithValue()：用于向 context 添加键值对，注意添加方式是以链状的方法

实现类3：timerCtx
timeout 机制
WithDeadline()： 可取消、有倒计时
WithTimeout()：一定程度上的重载，详见源码

[应用场景]
- web 编程中，一个请求对应多个 goroutine 之间的数据交互
- 超时控制
- 上下文控制

[其他]
- concurrent safe
- cannot create context from nil parent
- cancel func can call multiple time
*/
