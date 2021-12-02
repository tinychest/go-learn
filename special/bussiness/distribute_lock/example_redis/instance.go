package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

/*
引入依赖 github.com/gomodule/redigo v1.8.5
*/

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				return nil, err
			}
			// 没有密码，就不要进行这一步了
			// if _, err = c.Do("AUTH"); err != nil {
			// 	_ = c.Close()
			// 	return nil, err
			// }
			if _, err = c.Do("SELECT", 0); err != nil {
				_ = c.Close()
				return nil, err
			}
			return c, nil
		},
		// Other pool configuration not shown in this example.
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func incr() {
	client := pool.Get()
	defer client.Close()

	var lockKey = "counter_lock"
	var counterKey = "counter"

	// lock
	lockSuccess, err := redis.Bool(client.Do("SETNX", lockKey, 1))
	if err != nil || !lockSuccess {
		fmt.Println(err, "lock result: ", lockSuccess)
		return
	}

	// counter ++
	cntValue, err := redis.Int64(client.Do("GET", counterKey))
	if err == nil {
		cntValue++
		_, err := client.Do("SET", counterKey, cntValue)
		if err != nil {
			println("set value error!")
		}
	}
	println("current counter is ", cntValue)

	// unlock
	unlockSuccess, err := redis.Bool(client.Do("DEL", lockKey))
	if err == nil && unlockSuccess {
		println("unlock success!")
	} else {
		println("unlock failed", err)
	}
}

func main() {
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
