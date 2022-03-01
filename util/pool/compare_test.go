package pool

import (
	"context"
	"github.com/go-pkgz/syncs"
	"github.com/panjf2000/ants/v2"
	"go-learn/util"
	"runtime"
	"testing"
)

const (
	taskSum     = 1000000 // 计算次数
	taskQuality = 1000  // 单次计算阶乘中的 n
)

func TestSimple(t *testing.T) {
	directCase(t)
	poolCase(t)
	antsPoolCase(t)
	// tpSyncsCase(t)
}

func directCase(t *testing.T) {
	res := make([]int, 0, taskSum)

	t.Log("direct")
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				data, _ := forFactorial()
				res = append(res, data.(int))
			}
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

	t.Log("ants")
	util.TimeCost(
		func() {
			p, _ := ants.NewPool(100)
			for i := 0; i < taskSum; i++ {
				_ = p.Submit(func() {
					data, _ := forFactorial()
					res = append(res, data.(int))
				})

				p.Release()
			}
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
