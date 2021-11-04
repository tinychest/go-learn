package number

import (
	"fmt"
	"go-learn/util"
	"strconv"
	"testing"
)

// 小数位处理（还可以参见 util/format.go）
func TestDecimal(t *testing.T) {
	mathTest()
	fmtRoundTest()
}

func mathTest() {
	// Ceil（向上取整） 看后边所有
	fmt.Println(util.Ceil(0.070001, 4))
	// round（四舍五入） 只看指定位数的后面一位
	fmt.Println(util.Round(0.070009, 4))
	// floor（向下取整） 直接扔掉
	fmt.Println(util.Floor(0.07005, 4))
}

func fmtRoundTest() {
	// 四舍五入1
	i, _ := strconv.Atoi(fmt.Sprintf("%.f", 1.4))
	println(i)
	i, _ = strconv.Atoi(fmt.Sprintf("%.f", 1.5))
	println(i)
	i, _ = strconv.Atoi(fmt.Sprintf("%.f", 1.6))
	println(i)
}
