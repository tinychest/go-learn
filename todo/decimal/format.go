package decimal

import (
	"math"
)

func Ceil(f float64, b int) float64 {
	return math.Ceil(f*math.Pow10(b)) / math.Pow10(b)
}

func Round(f float64, b int) float64 {
	return math.Round(f*math.Pow10(b)) / math.Pow10(b)
}

func Floor(f float64, b int) float64 {
	return math.Floor(f*math.Pow10(b)) / math.Pow10(b)
}
