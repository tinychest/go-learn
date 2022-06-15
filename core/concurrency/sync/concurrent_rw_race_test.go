package sync

import (
	"sync"
	"testing"
)

// 来源：https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651451907&idx=1&sn=ec4f4fd7e033243fce4c42973d585744&chksm=80bb2cf1b7cca5e7331408439a4c7977faa9bc7970e740837b251b1e210c6a57000d51d5548f&scene=21#wechat_redirect
// 真正来源：https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549

// go test -v --race -run ^\QTestConcurrentRW\E$
// 以下样例都能被 --race 参数，检测到 DATA RACE 的
func TestConcurrentRW(t *testing.T) {
	// case1Test(t)
	// case2Test(t)
	case3Test(t)
}

// [可以看到现象]
// 原因：隐含的多段读
// 简单分析可以得到，就是 打印的读操作是遍历 a 的每一个元素，当 a 变了，相当于打到另外一个切片的元素了
// 应该加上读锁，我正在读取，还没有读取完的时候，不能进行写操作修改
func case1Test(t *testing.T) {
	type Config struct {
		a []int
	}
	cfg := &Config{}

	// 写
	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()

	// 读
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

// [可以看到现象]
// 多端写
func case2Test(t *testing.T) {
	s := make([]int, 5)

	go func() {
		i := 0
		for {
			i++
			s[0] = i
			s[1] = i + 1
			s[2] = i + 2
			s[3] = i + 3
			s[4] = i + 4
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", s)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// [看不到现象]
func case3Test(t *testing.T) {
	var a []int

	go func() {
		i := 0
		for {
			i++
			a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5, i + 6, i + 7, i + 8, i + 9}
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				t.Logf("%#v\n", a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
