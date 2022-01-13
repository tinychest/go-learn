package number

import (
	"fmt"
	"strconv"
	"testing"
)

const sum = 0.1

// 处理数位计算，在 Java 中是使用 BigDecimal，从本质说，从计算机组成原理说，计算机中的二进制，就不可能等价的表示出一些小数，Java 的针对这个做的处理是，当在一定精度内相等，就认为是相等
func TestPrecision(t *testing.T) {
	t.Log(1e-1 == sum)
	t.Log(0.8999999999+1e-9 >= 0.9)
	// 编译器提示 false，但是实际执行结果为 true
	t.Log(0.1+0.2 == 0.3)

	t.Log(CentToYuan(0) + "元")
	t.Log(CentToYuan(1) + "元")
	t.Log(CentToYuan(30) + "元")
	t.Log(CentToYuan(31) + "元")
	t.Log(CentToYuan(100) + "元")
	t.Log(CentToYuan(130) + "元")
	t.Log(CentToYuan(13000) + "元")
}

// 既然，这是计算机的根本，引出的问题，所以最好解决方式是，尽可能去避免这个，尽量使用 string 格式，去处理成希望的数据格式
func CentToYuan(cent int) (s string) {
	switch {
	case cent < 0:
		panic("invalid argument to CentToYuan")
	case cent == 0:
		return "0"
	}

	if cent < 100 {
		s = fmt.Sprintf("%03d", cent)
	} else {
		s = strconv.Itoa(cent)
		// s = fmt.Sprintf("%d", cent)
	}
	return s[:len(s)-2] + "." + s[len(s)-2:]
}