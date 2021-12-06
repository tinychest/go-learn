package pool

import (
	"github.com/panjf2000/ants/v2"
	"go-learn/util"
	"log"
	"testing"
)

const (
	taskQuality = 1000
	taskSum     = 1000
)

func TestSimple(t *testing.T) {
	res1 := make([]int, 0, taskSum)
	res2 := make([]int, 0, taskSum)
	res3 := make([]int, 0, taskSum)

	// 直接运算
	t.Log("direct")
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				data, _ := forFactorial()
				res1 = append(res1, data.(int))
			}
		})

	// 使用自定义任务池
	getter := func(index int) Task {
		return forFactorial
	}
	collector := func(data interface{}) {
		res1 = append(res2, data.(int))
	}

	t.Log("custom")
	util.TimeCost(
		func() {
			NewTaskCenter(taskSum, getter).Start(collector)
		})

	// 使用 ant
	t.Log("ants")
	p, _ := ants.NewPool(100)
	util.TimeCost(
		func() {
			for i := 0; i < taskSum; i++ {
				err := p.Submit(func() {
					data, _ := forFactorial()
					res3 = append(res3, data.(int))
				})
				if err != nil {
					log.Fatal(err)
				}
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
