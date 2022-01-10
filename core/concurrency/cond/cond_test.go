package cond

import (
	"sync"
	"testing"
)

/*
【sync.Cond】
[概念]
用于多个协程间的通知（概念同 Java 的同步操作）

（房间 = 锁）
- 房间里边同一时间只能有一个醒着的人

- c.l.Lock() 进入房间（自动门会锁上）
- c.l.Unlock() 走出房间（门锁会打开）
- c.Wait() 在房间里睡觉（持有锁的人睡觉了，自动门会打开）
- c.Signal() 敲醒房间里边睡觉的一个人（在房间里边，外边都可以敲醒）
- c.Broadcast() 敲醒房间里边睡觉的所有人（在房间里边，外边都可以敲醒）

[真实的业务场景]
多个 Goroutine 等待一个条件成立，然后，继续执行
*/

func TestCond(t *testing.T) {
	l := sync.Mutex{}
	c := sync.NewCond(&l)

	resource := 0

	// 生产者
	go func() {
		c.L.Lock()
		defer c.L.Unlock()

		for i := 0; i < 3; i++ {
			t.Log("生产者生产一个资源")
			resource++
		}
		c.Broadcast()
	}()

	// 多个消费者
	for i := 0; i < 3; i++ {
		go func(i int) {
			c.L.Lock()
			defer c.L.Unlock()

			t.Logf("消费者 %d 获取到锁", i)
			if resource <= 0 {
				t.Logf("没有资源，消费者 %d 等待", i)
				c.Wait() // 释放锁并挂起
				t.Logf("消费者 %d 被唤醒", i)
			}

			t.Logf("消费者 %d 消费一个资源", i)
			resource--
		}(i)
	}

	select {}
}
