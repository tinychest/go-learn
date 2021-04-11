package util

import (
	"math"
	"time"
)

// floatSum：浮点数
// keepDigitBit：要保留的小数位
func TrimDigitBit(floatSum float64, digitBit int) float64 {
	return math.Round(floatSum*math.Pow10(digitBit)) / math.Pow10(digitBit)
}

func ParseTime(tStr string) time.Time {
	if t, err := time.Parse(DateTimeFormat, tStr); err != nil {
		panic(err)
	} else {
		return t
	}
}

func FormatTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}
