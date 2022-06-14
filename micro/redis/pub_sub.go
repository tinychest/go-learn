package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// 使用 redis 来实现一个简单的消息队列机制

const chanName = "ccc"

func main() {
	conn := pool.Get()

	psc := redis.PubSubConn{Conn: conn}

	if err := psc.Subscribe(redis.Args{}.AddFlat(chanName)...); err != nil {
		panic(err)
	}
	done := make(chan error, 1)
	// start a new goroutine to receive message
	go func() {
		defer psc.Close()
		for {
			switch msg := psc.Receive().(type) {
			// 发生错误
			case error:
				fmt.Println("error", msg)
				done <- fmt.Errorf("redis pubsub receive err: %v", msg)
				return
			// publish（普通的事件发布）
			case redis.Message:
				fmt.Println("message", msg.Channel, string(msg.Data))
			// subscribe, unsubscribe, psubscribe or punsubscribe（订阅新增或减少 - 包括自己订阅的事件）
			case redis.Subscription:
				fmt.Println("subscription", msg.Channel, msg.Count, msg.Kind)
				// channel do not have any subscribed channel
				if msg.Count == 0 {
					done <- nil
					return
				}
			default:
				panic(fmt.Sprintf("unknown msg type: %s", msg))
			}
		}
	}()

	// GG 不支持这个操作：unexpected response line (possible server error or unsupported concurrent read by application)
	// sum, err := redis.Int(conn.Do("PUBSUB NUMPAT"))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("PUBSUM NUMPAT:", sum)

	<-done
	if err := psc.Unsubscribe(chanName); err != nil {
		panic(err)
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	// defer cancel()
	//
	// // health check
	// tick := time.NewTicker(time.Minute)
	// defer tick.Stop()
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		if err := psc.Unsubscribe(); err != nil {
	// 			panic(fmt.Errorf("redis pubsub unsubscribe err: %v", err))
	// 		}
	// 	case err := <-done:
	// 		return err
	// 	case <-tick.C:
	// 		if err := psc.Ping(""); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

}
