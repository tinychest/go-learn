package decimal

/*
一个比较好，且目前没有测出问题的，精度丢失补偿措施
*/

// FloorAccurate2 指定位数，向下取整
func FloorAccurate2(f float64, b int) float64 {
	f += AccuracyToleration
	return Floor(f, b)
}

// CeilAccurate2 指定位数，向上取整
func CeilAccurate2(f float64, b int) float64 {
	f -= AccuracyToleration
	return Ceil(f, b)
}

// RoundAccurate2 指定位数，四舍五入
func RoundAccurate2(f float64, b int) float64 {
	f += AccuracyToleration
	return Round(f, b)
}
