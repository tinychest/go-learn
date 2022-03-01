package decimal

import (
	"math"
	"testing"
)

func TestPrecision(t *testing.T) {
	// baseTest(t)
	// baseTest2(t)
	// baseTest3(t)
	// baseTest4(t)
	// baseTest5(t)

	// watchTest(t)
	// constTest(t)

	// classicDealTest(t)
}

// 0.1 + 0.2 != 0.3
// 该样例的结果和在 browser console 中执行 0.1 + 0.2 的结果一致
// 足以说明这和语言无关，是计算机底层原理导致的
func baseTest(t *testing.T) {
	a := 0.1
	b := 0.2

	t.Log(a + b)
	t.Log(a + b == 0.3)
}

// 上例是真实结果比预期大，这里再给出一个真实结果比预期结果小的
func baseTest2(t *testing.T) {
	a := 0.1
	b := 0.7

	t.Log(a + b)
}

// TODO 常量和变量带来的影响（关于常量浮点数是精度更高的数，需要给出官方文档地址）
func baseTest3(t *testing.T) {
	a := 0.03

	t.Log(0.03 + 0.005)
	t.Log(a + 0.005)
}

// TODO 类型转换带来的影响（定义变量和显示类型转换起到的效果一样）
func baseTest4(t *testing.T) {
	const a = 0.03

	t.Log(a + 0.005)
	t.Log(float64(a) + 0.005)
	t.Log(a == float64(a))
	t.Log(a + 0.005 == float64(a) + 0.005)
}

// TODO 计算带来的影响
func baseTest5(t *testing.T) {
	t.Log(0.03 + 0.005)
	t.Log(0.03 * float64(1) + 0.005)
}

// 再看看这个[狗头]
// 这是要否定这里所有的说法么，不不不，不是的，这里想说明的是：
// - 静态常量的精度是非常大的
// - Go 底层也有在一定误差就认为相等的
func constTest(t *testing.T) {
	const (
		a = 0.1
		b = 0.2
	)

	t.Log(a + b)
	t.Log(a + b == 0.3)
}

func watchTest(t *testing.T) {
	v := 0.00000007

	t.Logf("%.20f\n", v*1e2)
	t.Logf("%.20f\n", v*1e3)
	t.Logf("%.20f\n", v*1e4)
	t.Logf("%.20f\n", v*1e5)
	t.Logf("%.20f\n", v*1e6)
}

// 给出 精度容忍 的正确比较方法
// 这里的关键就是精度，就是你对精度容忍的度，你容忍的越大，结果就会越不准确
func classicDealTest(t *testing.T) {
	equalFunc := func(a, b float64) bool {
		return math.Abs(a - b) < 1e-10
	}
	a, b := 0.1, 0.2
	t.Log(equalFunc(a + b, 0.3))
}
