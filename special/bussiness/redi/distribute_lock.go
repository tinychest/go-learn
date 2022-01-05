package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

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
