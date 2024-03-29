package decimal

import (
	"fmt"
	"math"
	"strconv"
)

/*
doc.go 中给出的样例，告诉我们：
- 有限的小数位，只能近似的表示一个浮点数，所以，浮点数无法避免误差
- 我们应该思考得出满足当前上下文，满足业务的精度容忍值

- 精度定义的不好，或者经过逐步计算、转换，也会逐步放大了这个误差
- 所以实际计算应该采用高精度的数值类型去计算和处理，尽可能减少影响

下面给出容忍的实现（目前只支持正数）
例：假如业务中只考虑 4 位小数，那不凡考虑将精度容忍值定为 10^(-14)

从分析和测试样例上来说，这或许是一种非常好的最佳实践
之所以没有任何相关、类似的文章资料，可能是因为使用的限制吧
*/

const (
	TolerationBit = 10
	Toleration    = 1e-10
)

// ToleCeil 指定位数，向上取整
// 精度容忍：如果取整丢弃的部分足够大，才能向上进位
// 实际：如果你丢弃的部分足够大，那么我减去一个微不足道的值，应该也不会对你造成任何影响；如果你丢弃的部分是因为误差（必定很小），那么这个减去的值将会把你打回原形
// n ∈ (1, 2)、math.Ceil(n) → 2
// ✔ 小边界-低案例：1（0.99...999） → 1（1）
// ❌ 小边界-高案例：1（1.00...001） → 1（2）
// ✔ 大边界-低案例：2（1.99...999） → 2（2）
// ❌ 大边界-高案例：2（2.00...001） → 2（3）
//
// 误差点：站在进位分界线，以为不会进，实际进了
func ToleCeil(f float64, b int) float64 {
	check(f, b)

	b10 := math.Pow10(b)

	return math.Ceil((f-Toleration)*b10) / b10
}

// ToleRound 指定位数，四舍五入
// n ∈ [0.5, 1.5)、math.Round(n) → 1
// ❌ 小边界-低案例：0.5（0.499...999） → 1（0）
// ✔ 小边界-高案例：0.5（0.500...001） → 1（1）
// ❌ 大边界-低案例：1.5（1.49....999） → 2（1）
// ✔ 大边界-高案例：1.5（1.50...001） → 2（2）
//
// 误差点：站在进位分界线，以为可以进，实际没有
func ToleRound(f float64, b int) float64 {
	check(f, b)

	b10 := math.Pow10(b)

	return math.Round((f+Toleration)*b10) / b10
}

// ToleFloor 指定位数，向下取整
// 假如实际值是 "0.99999999999999"（调用者认为是 1，并期望得到 1），但是通过标准库的方式处理，将得到 0
// n ∈ [1, 2)、math.Floor(n) → 1
// ❌ 小边界-低案例：1（0.99....999） → 1（0）
// ✔ 小边界-高案例：1（1.00...001） → 1（1）
// ❌ 大边界-低案例：2（1.99....999） → 2（1）
// ✔ 大边界-高案例：2（2.00...001） → 2（2）
//
// 误差点：站在进位分界线，以为可以进，实际没有
func ToleFloor(f float64, b int) float64 {
	check(f, b)

	b10 := math.Pow10(b)

	return math.Floor((f+Toleration)*b10) / b10
}

// 总结
// - 实际上不用取大小边界，因为运算规则看重边界，并不关注哪个边界
// （补偿的值很小，肯定比业务场景涉及的小数范围小得多）
// - Ceil 误进位，需要“刁难”（理想精确的值进行“刁难”不会对结果产生改变）
// - Round、Floor 误退位，需要补偿（理想精确的值进行补偿不会对结果产生改变）

func check(f float64, b int) {
	if f < 0 {
		panic("param f can not be negative")
	}
	if b > TolerationBit {
		panic(fmt.Errorf("param b is too big, max support is %d", TolerationBit))
	}
}

func ToleCeilStr(f float64, b int) string {
	return strconv.FormatFloat(ToleCeil(f, b), 'f', b, 64)
}

func ToleRoundStr(f float64, b int) string {
	return strconv.FormatFloat(ToleRound(f, b), 'f', b, 64)
}

func ToleFloorStr(f float64, b int) string {
	return strconv.FormatFloat(ToleFloor(f, b), 'f', b, 64)
}
