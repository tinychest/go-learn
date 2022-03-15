package pool

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func TestCommonPool(t *testing.T) {
	// Use the common pool.
	defer ants.Release()

	runTimes := 1000

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	runTimes := 1000
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(1))
	}
	wg.Wait()

	t.Logf("running goroutines: %d\n", p.Running())
	t.Logf("finish all tasks, result is %d\n", sum)
}
