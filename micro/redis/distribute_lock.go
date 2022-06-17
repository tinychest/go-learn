package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// 首先还是要将锁的概念放到脑海里，有并发操作，操作同一资源，可以通过锁来解决并发问题
// go-zero 中有给出现成的实现方案
// [参考] https://juejin.cn/post/7041375517580689439
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
// 排他性、互斥性：Redis 的命令基本都是原子操作
// 防死锁：Redis 的 SETNX、SET NX 命令都有提供超时时间的参数
// 可重入：Redis 相同指令除了 key，还应该带上全局唯一的 value，来允许再次获取锁，同时刷新超时时间，防止重入时，因超时被释放
//     加锁和解锁都不能简单的执行一下 SET 带有相关参数的命令，而是通过 lua 脚本执行一段原子操作的逻辑
//    （实际就是，加锁、解锁 脚本需要进行一个对值的比对）
// 高性能、高可用：高性能就不说了；Redis 支持主从和集群以及哨兵，来应对高并发和单点故障等问题

const (
	lockKey          = "counter_lock"
	defaultExpireSec = 5

	lockTryInterval = 2 * time.Second // 加锁失败，进行重试时的，重试间隔时间
	lockTryTimes    = 3               // 加锁失败，最多重试的次数

	// 加锁脚本（可重入）
	// 判断 key 持有的 value 是否等于传入的 value
	// - 相等：说明是再次获取锁：更新时间，防止重入时过期
	// - 不相等：说明是第一次获取锁：通过 NX 参数去设置键值，也就是尝试获取一次锁
	lockCommand = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "ok"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "EX", ARGV[2])
end`

	// 解锁脚本（可重入）
	unlockCommand = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
)

func lock(key string, value interface{}) (bool, error) {
	c := pool.Get()
	defer c.Close()

	// 普通加锁（不可重入）
	// reply, err := c.Do("SET", lockKey, value, "NX", "EX", defaultExpireSec)
	// 加锁（可重入）
	res, err := redis.String(c.Do("EVAL", lockCommand, 1, key, value, defaultExpireSec))
	if err != nil {
		return false, err
	}
	return res == "OK", nil
}

// 自定义重试机制的 阻塞实现
func lockPre(key string, value interface{}) error {
	var timer *time.Timer

	for i := 0; i < lockTryTimes+1; i++ {
		if i != 0 {
			fmt.Printf("第 %d 尝试加锁\n", i)
		}
		ok, err := lock(key, value)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		if timer == nil {
			timer = time.NewTimer(lockTryInterval)
		} else {
			timer.Reset(lockTryInterval)
		}
		<-timer.C
	}

	return fmt.Errorf("尝试了 %d 次，仍然枷锁失败", lockTryTimes)
}

func unlock(key string, value interface{}) error {
	c := pool.Get()
	defer c.Close()

	res, err := redis.Int64(c.Do("EVAL", unlockCommand, 1, key, value))
	if err != nil {
		return err
	}
	if res != 1 {
		return fmt.Errorf("解锁失败 key:%s value:%d", key, value)
	}
	return nil
}
