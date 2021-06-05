package util

import (
	"go-learn/const/time_format"
	"math"
	"time"
)

// 保留指定的小数位 - 向上
func Ceil(s float64, b int) float64 {
	return math.Ceil(s*math.Pow10(b)) / math.Pow10(b)
}

// 保留指定的小数位 - 四舍五入
func Round(s float64, b int) float64 {
	return math.Round(s*math.Pow10(b)) / math.Pow10(b)
}

// 保留指定的小数位 - 向下（准确）
func Floor(s float64, b int) float64 {
	return math.Floor(s*math.Pow10(b)) / math.Pow10(b)
}

func Floor2(s float64, b int) float64 {
	return math.Round(s*math.Pow10(b))/math.Pow10(b)
}

// 参数格式必须是 2006-01-02 15:04:05
func ParseTime(s string) time.Time {
	if t, err := time.Parse(time_format.DateTimeFormat, s); err != nil {
		panic(err)
	} else {
		return t
	}
}

func FormatTime(t time.Time) string {
	return t.Format(time_format.DateTimeFormat)
}
