package number

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// 小数位处理（还可以参见 util/format.go）
// Ceil 看不到源码实现
// Sprintf 可以看到源码实现，就是单纯的处理 format
// Round 看得到源码实现，但是和 Sprintf 的实现不同
// Floor 看不到源码实现
func TestAk(t *testing.T) {
	// 向上取整
	println(int(math.Ceil(1.4)), int(math.Ceil(1.5)),int(math.Ceil(1.6)))

	// 四舍五入1
	i, _ := strconv.Atoi(fmt.Sprintf("%.f", 1.4))
	println(i)
	i, _ = strconv.Atoi(fmt.Sprintf("%.f", 1.5))
	println(i)
	i, _ = strconv.Atoi(fmt.Sprintf("%.f", 1.6))
	println(i)

	// 四舍五入2
	println(int(math.Round(1.4)), int(math.Round(1.5)), int(math.Round(1.6)))

	// 向下取整
	println(int(math.Floor(1.4)), int(math.Floor(1.5)), int(math.Floor(1.6)))
}
