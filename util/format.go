package util

import (
	"go-learn/const/time_format"
	"math"
	"time"
)

// Ceil 保留指定的小数位 - 向上
func Ceil(s float64, b int) float64 {
	return math.Ceil(s*math.Pow10(b)) / math.Pow10(b)
}

// Round 保留指定的小数位 - 四舍五入
func Round(s float64, b int) float64 {
	return math.Round(s*math.Pow10(b)) / math.Pow10(b)
}

// Floor 保留指定的小数位 - 向下（准确）
func Floor(s float64, b int) float64 {
	return math.Floor(s*math.Pow10(b)) / math.Pow10(b)
}

func Floor2(s float64, b int) float64 {
	return math.Round(s*math.Pow10(b)) / math.Pow10(b)
}

func ParseTime(s string) time.Time {
	return ParseFmtTime(s, time_format.FmtDateTime)
}

func ParseFmtTime(s string, format string) time.Time {
	if t, err := time.Parse(format, s); err != nil {
		panic(err)
	} else {
		return t
	}
}

func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

func Now() string {
	return FormatTime(time.Now(), time_format.FmtDateTime)
}

func NowCN() string {
	return FormatTime(time.Now(), time_format.FmtDateTimeCN)
}

func NowDate() string {
	return FormatTime(time.Now(), time_format.FmtDate)
}

func NowDateCN() string {
	return FormatTime(time.Now(), time_format.FmtDateCN)
}
