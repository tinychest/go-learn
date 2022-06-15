package atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	a []int
}

// go test -v --race -run \QTestValue\E
func TestValue(t *testing.T) {
	// valueMisTest(t)
	// valueMis2Test(t)
	// valueTest(t)
	// value2Test(t)
}

// atomic.Value 相当于没有使用，因为 atomic.Value 实例存储的数据的地址从未改变
func valueMisTest(t *testing.T) {
	cfg := &Config{}

	var v = new(atomic.Value)

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

// 虽然现象上没有错误，但是存在 DATA RACE
func valueMis2Test(t *testing.T) {
	cfg := &Config{}

	go func() {
		i := 0
		for {
			i++
			cfg = &Config{
				a: []int{i, i + 1, i + 2, i + 3, i + 4, i + 5},
			}
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", cfg)
			}
		}()
	}
	wg.Wait()
}

// 正确使用
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

func value2Test(t *testing.T) {
	var v atomic.Value

	go func() {
		i := 0
		for {
			i++
			v.Store([]int{i, i + 1, i + 2, i + 3, i + 4, i + 5})
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
