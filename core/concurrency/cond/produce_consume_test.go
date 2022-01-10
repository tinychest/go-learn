package cond

import (
	"sync"
	"testing"
	"time"
)

/*
生产组成员数 和 消费组成员数相等
实现消费组一来就休息，然后生产组每个成员一人生产一次，然后，生产队算生产完毕，喊消费组干活，消费组每个成员都消费一起，继续喊生产组起来干活，如此循环...
*/

func TestProduceAndConsume(t *testing.T) {
	const (
		prodSum = 3
		consSum = 3
	)

	l := &sync.Mutex{}
	prod := sync.NewCond(l)
	cons := sync.NewCond(l)

	resource := 0
	prodOk := false

	for i := 0; i < consSum; i++ {
		go func(i int) {
			for {
				l.Lock()

				t.Logf("消费者 %d 准备消费", i)
				if !prodOk {
					t.Logf("消费者 %d 发现没得消费 → 睡觉", i)
					cons.Wait()
					t.Logf("消费者 %d 被唤醒", i)
				}
				t.Logf("消费者 %d 开始消费", i)
				time.Sleep(time.Second)
				resource--
				if resource == 0 {
					prodOk = false
					t.Logf("消费者 %d 发现消费完了，唤醒生产组", i)
					prod.Broadcast()
				}

				t.Logf("消费者 %d 消费完了，睡觉", i)
				cons.Wait()
				t.Logf("生产者 %d 被唤醒", i)

				l.Unlock()
			}
		}(i)
	}

	time.Sleep(time.Second)

	for i := 0; i < prodSum; i++ {
		go func(i int) {
			for {
				l.Lock()

				t.Logf("生产者 %d 准备生产", i)
				if prodOk {
					t.Logf("生产者 %d 发现生产满了，睡觉", i)
					prod.Wait()
					t.Logf("生产者 %d 被唤醒", i)
				}
				t.Logf("生产者 %d 开始生产", i)
				time.Sleep(time.Second)
				resource++
				if resource == prodSum {
					prodOk = true
					t.Logf("生产者 %d 发现生产完了，唤醒消费组", i)
					cons.Broadcast()
				}

				t.Logf("生产者 %d 生产完了，睡觉", i)
				prod.Wait()
				t.Logf("生产者 %d 被唤醒", i)

				l.Unlock()
			}
		}(i)
	}

	select {}
}
