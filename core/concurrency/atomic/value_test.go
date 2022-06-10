package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// 来源：https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651451907&idx=1&sn=ec4f4fd7e033243fce4c42973d585744&chksm=80bb2cf1b7cca5e7331408439a4c7977faa9bc7970e740837b251b1e210c6a57000d51d5548f&scene=21#wechat_redirect
// 真正来源：https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549

type Config struct {
	a []int
}

func TestValue(t *testing.T) {
	// normalTest(t)
	// valueMisTest(t)
	valueTest(t)
	// syncTest(t)
}

// TODO 能得到数字不连续的现象，但是将 a 改成单个切片，就看不到这个现象了
func normalTest(t *testing.T) {
	cfg := &Config{}

	// 一个 goroutine 不停写
	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()

	// 多个 goroutine 读
	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// TODO why
func valueMisTest(t *testing.T) {
	cfg := &Config{}

	var v atomic.Value

	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", v.Load())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func valueTest(t *testing.T) {
	cfg := &Config{}

	var v atomic.Value

	go func() {
		i := 0
		for {
			i++
			cfg = &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", v.Load())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func syncTest(t *testing.T) {
	cfg := &Config{}

	var mux sync.RWMutex

	go func() {
		i := 0
		for {
			i++
			mux.Lock()
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			mux.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				mux.RLock()
				t.Logf("%#v\n", cfg)
				mux.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
