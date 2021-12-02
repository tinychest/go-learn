package decimal

import (
	"fmt"
	"math"
)

/*
format.go 中定义的 Ceil、Round、Floor 的方法逻辑有问题么，没有
但是，因为计算机底层的二进制模拟，导致了方法实际的参数和调研者预期的不一样，所以调用者想当然认为是方法的逻辑实现问题
也就是调用者希望方法能够处理差异，给出调用者预期的答案，本质上也就是 “精度容忍”（其他高级语言，无一例外都是这样做的）

下面给出容忍的实现（目前只支持正数）

实际测试你会发现，精度定义的不好，或者因为小的精度差异，经过逐步计算，放大了这个差异
所以实际计算应该采用高精度的数值类型去计算和处理，尽可能减少精度累计带来的影响
*/

const (
	AccuracyToleration = 1e-10
	AllowMaxBit        = 10
)

// FloorAccurate 指定位数，向下取整
// 如果取整丢弃的部分过大，就应该判定是因为精度，导致的过大丢弃部分
// 假如实际值是 "0.99999999999999"（调用者认为是 1，并期望得到 1），但是通过原生的方式处理，将得到 0
func FloorAccurate(f float64, b int) float64 {
	check(f, b)

	fb := f * math.Pow10(b)
	i := int(fb)
	if 1-(fb-float64(i)) < AccuracyToleration {
		i += 1
	}
	return float64(i) / math.Pow10(b)
}

// CeilAccurate 指定位数，向上取整
// 如果取整丢弃的部分足够大，才能向上进位
// 假如实际值是 "1.000000000000001"（调用者认为是 1，并期望得到 1），但是通过原生的方式处理，将得到 2
func CeilAccurate(f float64, b int) float64 {
	check(f, b)

	fb := f * math.Pow10(b)
	i := int(fb)
	if fb-float64(i) > AccuracyToleration {
		i += 1
	}
	return float64(i) / math.Pow10(b)
}

// RoundAccurate 指定位数，四舍五入
func RoundAccurate(f float64, b int) float64 {
	check(f, b)

	b1 := f * math.Pow10(b)
	i1 := int(b1)
	b2 := b1 + 0.5
	i2 := int(b2)

	res := i1

	if 1-(b2-float64(i2)) < AccuracyToleration {
		res += 1
	}
	return float64(res) / math.Pow10(b)
}

func check(f float64, b int) {
	if f < 0 {
		panic("param f can not be negative")
	}
	if b > AllowMaxBit {
		panic(fmt.Errorf("param b is too big, max support is %d", AllowMaxBit))
	}
}
