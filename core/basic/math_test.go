package basic

import (
	"fmt"
	"math"
	"testing"
)

// 在 java 中进行除以 0 的计算，将会造成 Exception
// Go 如果是同概念的话，应该会 panic，但实际是返回 Inf

// 0 / 0 → NaN
// + / 0 → +Inf
// - / 0 → -Inf
func TestDivisionZero(t *testing.T) {
	zero := 0.0

	nan := 0 / zero
	tInf := 1 / zero
	nInf := -1 / zero

	inf := math.Inf(1)
	if 0x7FF0000000000000 < inf {
		fmt.Println("小")
	}
	if 0x7FF0000000000000 > inf {
		fmt.Println("大")
	}

	fmt.Println(nan)               // NaN
	fmt.Println(math.NaN())        // NaN
	fmt.Println(nan == math.NaN()) // TODO false???
	// Inf 参数 >= 0 返回正无穷
	fmt.Println(tInf == math.Inf(1))
	// Inf 参数 <= 0 返回负无穷
	fmt.Println(nInf == math.Inf(-1))
}
