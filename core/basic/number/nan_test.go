package number

import (
	"fmt"
	"math"
	"testing"
)

// 分母为 0 的除法结果是 nan，但是 Go 里边有做处理任何和 nan 的比较操作返回的都是 false
func TestNanCompare(t *testing.T) {
	f := 0.0
	nan := 0 / f

	inf := math.Inf(1)
	if 0x7FF0000000000000 < inf {
		fmt.Println("小")
	}
	if 0x7FF0000000000000 > inf {
		fmt.Println("大")
	}

	fmt.Println(nan)               // NaN
	fmt.Println(math.NaN())        // NaN
	fmt.Println(nan == math.NaN()) // false
	fmt.Println(math.NaN() == math.NaN()) // false
}
