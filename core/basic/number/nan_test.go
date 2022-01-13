package number

import (
	"math"
	"testing"
)

// 分母为 0 的除法结果是 nan，但是 Go 里边有做处理任何和 nan 的比较操作返回的都是 false
func TestNanCompare(t *testing.T) {
	f := 0.0
	nan := 0 / f

	inf := math.Inf(1)
	if 0x7FF0000000000000 < inf {
		t.Log("小")
	}
	if 0x7FF0000000000000 > inf {
		t.Log("大")
	}

	t.Log(nan)                      // NaN
	t.Log(math.NaN())               // NaN
	t.Log(nan == math.NaN())        // false
	t.Log(math.NaN() == math.NaN()) // false
}
