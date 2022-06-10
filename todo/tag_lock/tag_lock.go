package tag_lock

// 标识锁的核心（并发编程）
// 只要向着将资源作为锁对象的思路，很能很好的瓦解和突破，特定的标识要能获取到特定的锁（同一个）
// - 存锁和上锁分别对应读和写，所以需要使用并发安全的 map（同时借助 LoadOrStore 的原语操作来防止锁覆盖）
// - 锁肯定会释放的，不用担心这个（defer unlock）
//
// 实际业务中
// - 最好给锁设定超时时间（当然，不是必须，本就是因为过于消耗资源，即，消耗的时间较长，才设定的锁机制）
// - 假定某一次的锁时间特别长，并且不断有新的请求到来，这样会导致不断累积请求（Goroutine），最终导致服务宕机，所以应该设置一个令牌桶，获取到令牌，才能服务请求（相当于加入等待队列）否则直接拒绝
//
// 标识锁分 标识等待 和 标识存在 两种实现，意思分别是
// - 标识等待 是指，根据特定标识去查看是否已经别占用，如果占用就等待（阻塞式方法调用）
// - 标识存在 是指，根据特定标识去查看是否已经别占用，会立即返回是否占用的结果（非阻塞式方法调用）

// SyncLocker 同步标识锁定义
type SyncLocker interface {
	Lock(key interface{})
	Unlock(key interface{}) error
}

type StatusLocker interface {
	Lock(key interface{}) bool
	Unlock(key interface{}) error
}
