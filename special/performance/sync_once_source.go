package performance

/*
自己的疑问，这个博主说的贼清楚：
https://blog.csdn.net/grace_yi/article/details/103621450

Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.
To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages.

1、atomic.LoadUnit32 方法的作用是可见性？
答案是的，是为了保证可见性，并且就在你看过的 Go 官方的内存模型一文中说到了

2、为什么 doSlow 方法中直接使用 o.done == 0 而没有用原子操作？
因为在临界区中，被锁上了，临界区中能够保证可见

3、为什么修改 o.done 的值用了原子操作，不是说在临界区么？
是呀，可是临界区外还有访问这个变量的啊，原子操作一定是配对的读写才有用
解析起来就和计算机底层原理相关了，每颗核都有自己的高速缓存，加上原子操作，就意味着无论读取都应该经过公共缓存

4、为什么不用 unit8 或者 bool？
额，Go 没有提供

5、提点：hot path，详见 performance\sync_once_test.go

6、为什么源码中提到使用自旋锁实现 Once 不行？
源码写的很清楚，只是英文有点难理解，这里解释一下：
如果使用了 cas 自旋锁的实现，当多个 goroutine 在竞争 cas 锁的时候，只要有一个竞争到，那么这个就是执行里边的逻辑，而其他的 goroutine 将直接返回
没错：直接返回，如果 Once 真的采用 cas 来设计，那将变的及其冗余无用，根本对不上大多数的应用场景，如懒汉单例
*/
