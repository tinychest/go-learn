package decimal

import (
	"math"
)

// 这里定义的方法，是为指定的浮点数按照指定策略保留指定小数位数的数值处理方法
// 但是没有经过任何处理，不能使用，存在问题

// BadCeil 保留指定的小数位 - 向上
// math.Ceil returns the greatest integer value less than or equal to x.
// 注意：看有没有小数
func BadCeil(s float64, b int) float64 {
	bp := math.Pow10(b)
	return math.Ceil(s*bp) / bp
}

// BadRound 保留指定的小数位 - 四舍五入
// 注意：只看后边的一位小数
func BadRound(s float64, b int) float64 {
	bp := math.Pow10(b)
	return math.Round(s*bp) / bp
}

// BadFloor 保留指定的小数位 - 向下
// math.Floor returns the greatest integer value less than or equal to x.
// 注意：看有没有小数
func BadFloor(s float64, b int) float64 {
	bp := math.Pow10(b)
	return math.Floor(s*bp) / bp
}