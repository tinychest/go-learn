package todo

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestQ2(t *testing.T) {
	strMap := map[int]string{1: "a", 2: "b", 3: "c"}

	strMap[4] = "d"

	delete(strMap, 0)

	fmt.Printf("%v", strMap)
}

// 为什么没有赋值成功，需要复习一下 defer 的知识点（确定了堆内存中的地址，你修改这个记录地址的变量又有什么用呢）
func TestQ3(t *testing.T) {
	fmt.Println(test())
}

func test() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()

	panic("gg")
	return err
}

func TestQ4(t *testing.T) {
	var a uint = 0
	var b uint = 1
	c := a - b

	fmt.Println(c == math.MaxUint)

	// >= 63 的结果都是一样的
	// m := uint(math.Pow(2, 62) - 1)
	// fmt.Println(m)
}

// a 是切片 for 中修改会影响接下来遍历的元素
// a 是数组 for 中修改不会影响接下来遍历的元素
func TestQ5(t *testing.T) {
	a := []int{1, 2, 3} // 切片 101, 300, 103
	// a := [3]int{1, 2, 3} // 数组 101 102 103
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
		}
		a[k] = 100 + v
	}

	fmt.Println(a)
}
