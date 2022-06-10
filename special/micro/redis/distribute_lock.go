package main

import (
	"github.com/gomodule/redigo/redis"
	"sync"
)

// 首先还是要将锁的概念放到脑海里，有并发操作，操作同一资源，可以通过锁来解决并发问题
//
// [应用场景]
// 在单机服务实例中通过一个 sync.Mutex 修饰特定的方法，这是简单的单机锁
// > 假定 MySQL 每秒能够应对 2k 个请求，而 Redis 每秒能够应对 10W 个请求
// > 其实就是 MySQL 在某些场景下来做分布式锁性能不行，当然，什么事都不是绝对的
// 如果，现在有需求，为了服务高可用，或是有了高并发的场景，做特定事情的服务有了多个，各自有者各自的锁，相当于没有锁
// 也就是锁需要由本地进程换个位置，放到一个所有实例都可以访问到的地方，很明显 MySQL 扛不住，选用 Redis 会发现是最合适的，成本小，能够轻易抗住并发量
// - 秒杀（锁住库存 - 防止超卖）
// - 商品下单（锁住订单号 - 防止重复下单）
// - 用户余额（锁住用户余额 - 假如有多个服务实例可以操作用户余额，想要用户余额不会出错，就应该使用锁）
//
// [特性]
// 排他性（互斥性）、防死锁、可重入、高性能
//
// 效率性：避免不同节点重复相同的工作
// 正确性：避免破坏正确性的发生，两个节点同时操作同一条数据
// 支持阻塞和非阻塞式：
// 支持公平和非公平
//
// [实现方式]
// MySQL、ZK、Redis
// 拿 Redis 来说
//
// 互斥性：Redis 提供了很多原子性操作命令，可以用于保证该点
// 防死锁：Redis 的 SET、SETNX 命令都有提供有效时间的参数，可以保证
// 阻塞和非阻塞：TODO 需要在代码实现上做文章
// 可重入：TODO 基于 Redis 的分布式锁，如果体现出重入性

const (
	lockKey    = "counter_lock"
	counterKey = "counter"
)

func incr() {
	c := pool.Get()
	defer c.Close()

	// lock
	ok, err := lock(lockKey)
	if err != nil {
		panic(err)
	}
	if !ok {
		// 锁被占用；可以自旋，每隔一段时间自旋一次，自旋一定次数后，认定失败
		// 注：你业务不能说，这个怎么还能失败。而是应该考虑如何尽可能减少失败的可能性，以及失败该如何处理
	}

	// 业务操作（这里是为 Redis 的一个值自增 1）
	cntValue, err := redis.Int64(c.Do("GET", counterKey))
	if err == nil {
		cntValue++
		_, err := c.Do("SET", counterKey, cntValue)
		if err != nil {
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// unlock
	ok, err = unlock(lockKey)
	if err != nil {
		panic(err)
	}
	if !ok {
		// 解锁失败要完蛋
		panic("unexpected unlock fail")
	}
}

func lock(key string) (bool, error) {
	c := pool.Get()
	defer c.Close()

	ok, err := redis.Bool(c.Do("SETNX", key, 1))
	if err != nil {
		return false, err
	}
	return ok, nil
}

func unlock(key string) (bool, error) {
	c := pool.Get()
	defer c.Close()

	ok, err := redis.Bool(c.Do("DEL", key))
	if err == nil {
		return false, err
	}
	return ok, nil
}

func main1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}
	wg.Wait()
}
