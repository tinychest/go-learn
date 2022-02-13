package pool

import (
	"github.com/panjf2000/ants/v2"
	"go-learn/util"
	"testing"
)

const (
	taskSum     = 10000 // 计算次数
	taskQuality = 1000  // 单次计算阶乘中的 n
)

func TestSimple(t *testing.T) {
	directCase(t)
	poolCase(t)
	antPoolCase(t)
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

	getter := func(index int) Task {
		return forFactorial
	}
	collector := func(data interface{}) {
		res = append(res, data.(int))
	}
	c := NewCenter(taskSum, getter)

	t.Log("custom")
	util.TimeCost(
		func() {
			c.Start(collector, 0)
			if err := c.Error(); err != nil {
				t.Log(err)
			}
		})
}

func antPoolCase(t *testing.T) {
	res := make([]int, 0, taskSum)

	p, _ := ants.NewPool(100)

	t.Log("ants")
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				_ = p.Submit(func() {
					data, _ := forFactorial()
					res = append(res, data.(int))
				})

				p.Release()
			}

			// p.Cap()
			// p.Tune(1)
			//
			// p.Running()
			// p.Free()
			//
			// p.Reboot()
			// p.Release()
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
