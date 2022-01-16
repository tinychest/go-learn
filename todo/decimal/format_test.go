package decimal

import (
	"math"
	"math/big"
	"testing"
)

/*
通过观察下面的样例的结果，你可能会得出常量的值总是准确的，问题都是出在变量上的结论
这个结论太肤浅了，实际上，Go 会对常量进行更高精度的存储和处理
*/

func TestBasic(t *testing.T) {
	const c1 = 0.03
	var v1 = 0.03

	f := func(a, b float64) {
		if a == b {
			t.Log(true)
			return
		}
		t.Log(false, a, b)
	}

	f(c1, v1)
	f(c1*1.0, v1*1.0)
	f(c1*1+0.005, v1*1+0.005)
}

func TestPBasic(t *testing.T) {
	const c1 = 0.03
	var v1 = c1

	t.Log(c1*1 + 0.005)
	t.Log(v1*1 + 0.005)

	// 好像并没有什么用
	b1 := big.NewFloat(c1)
	b1.Mul(b1, big.NewFloat(1))
	b1.Add(b1, big.NewFloat(0.005))
	t.Log(b1.Float64())
}

func TestRound(t *testing.T) {
	// 【缩小的案例】
	// [0.035] 期望结果 0.040000
	t.Logf("%f", Round(0.03+0.005, 2))
	// [0.034999999999999996] 期望 0.04 结果 0.030000
	t.Logf("%f", Round(0.03*float64(1)+0.005, 2))

	// 【放大的案例】
	const v1 = 0.07
	v2 := v1

	// [7] 期望结果 7.000000
	t.Logf("%f", Round(v1*100, 0))
	// [7.000000000000001] 期望结果 7.000000
	t.Logf("%f", Round(v2*100, 0))
}

func TestCeil(t *testing.T) {
	v := 0.07

	// 错误答案
	t.Logf("%f", math.Ceil(v*100))
	// 错误答案
	t.Logf("%f", Ceil(v*100, 0))
	// 正确答案
	t.Logf("%f", CeilAccurate(v*100, 0))
}

func TestFloor(t *testing.T) {
	t.Logf("%f", math.Floor((0.03*float64(1)+0.005)*math.Pow10(3)))
	t.Logf("%f", Floor(0.03*float64(1)+0.005, 21))
	t.Logf("%f", FloorAccurate(0.03*float64(1)+0.005, 21))
}

func TestName(t *testing.T) {
	v := 0.00000007

	// t.Log(v * 100)     // 7.000000000000001e-06
	// t.Log(v * 1000)    // 7.000000000000001e-05
	// t.Log(v * 10000)   // 0.0007000000000000001
	// t.Log(v * 100000)  // 0.007
	// t.Log(v * 1000000) // 0.07

	// t.Logf("%.20f\n", v*100)
	// t.Logf("%.20f\n", v*1000)
	// t.Logf("%.20f\n", v*10000)
	// t.Logf("%.20f\n", v*100000)
	// t.Logf("%.20f\n", v*1000000)

	v2 := v * 10000
	v3 := v2 * 10000000
	t.Log(v3)
}

/*
【主题】
击溃下面的实现

【Step01】
首先标准库 Ceil 的问题在于，将多出的值进 1 了 → 只能从实际数比预期大出发

只要这个 减操作 中的 忽略精度 解决不了实际因精度偏大的数据就行了，即放大误差

但是实际，上面的案例，希望放大误差，但是当达到一定的体量，误差就没了（这可是 float64，精度误差怎么可能大到这种程度）

【Step02】
没有了，想不出来
*/
func cc(s float64, b int) float64 {
	s -= 1e-10
	return math.Ceil(s*math.Pow10(b)) / math.Pow10(b)
}
