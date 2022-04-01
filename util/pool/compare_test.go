package pool

import (
	"context"
	"github.com/go-pkgz/syncs"
	"github.com/panjf2000/ants/v2"
	"go-learn/util"
	"runtime"
	"sync"
	"testing"
)

// TODO 自定义并发机制
//   每个任务等待被执行有一个超时时间
//   等待每个任务执行完成有一个超时时间

// TODO ants 执行的任务 panic 了，就 panic 了，不考虑中断继续发放任务（默认池子的数量现象是这样）
//   自己的池子就是因为考虑了这个，整的特别复杂

const (
	taskSum     = 100000 // 计算次数
	taskQuality = 1000  // 单次计算阶乘中的 n
)

func TestSimple(t *testing.T) {
	directCase(t)
	directGoCase(t)
	poolCase(t)
	antsPoolCase(t)
	// tpSyncsCase(t)
}

func directCase(t *testing.T) {
	t.Log("direct")
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				_, _ = forFactorial()
			}
		})
}

func directGoCase(t *testing.T) {
	var wg sync.WaitGroup

	t.Log("direct go")
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				wg.Add(1)
				go func() {
					_, _ = forFactorial()
					wg.Done()
				}()
			}
			wg.Wait()
		})
}

func poolCase(t *testing.T) {
	res := make([]int, 0, taskSum)

	t.Log("custom")
	util.TimeCost(
		func() {
			getter := func(index int) Task {
				return forFactorial
			}
			collector := func(data interface{}) {
				res = append(res, data.(int))
			}
			c := NewCenter(taskSum, getter)
			c.Start(collector, 0)
			if err := c.Error(); err != nil {
				t.Log(err)
			}
		})
}

func antsPoolCase(t *testing.T) {
	res := make([]int, 0, taskSum)
	var wg sync.WaitGroup

	t.Log("ants")
	util.TimeCost(
		func() {
			p, _ := ants.NewPool(100)
			defer p.Release()

			for i := 0; i < taskSum; i++ {
				wg.Add(1)
				_ = p.Submit(func() {
					data, _ := forFactorial()
					res = append(res, data.(int))
					wg.Done()
				})
			}
			wg.Wait()
		})
}

func tpSyncsCase(t *testing.T) {
	res := make([]int, 0, taskSum)

	t.Log("syncs")
	util.TimeCost(func() {
		swg := syncs.NewSizedGroup(runtime.GOMAXPROCS(8))
		for i := 0; i < taskSum; i++ {
			swg.Go(func(ctx context.Context){
				data, _ := forFactorial()
				res = append(res, data.(int))
			})
		}
		swg.Wait()
	})
}

func forFactorial() (interface{}, error) {
	return factorial(taskQuality), nil
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
