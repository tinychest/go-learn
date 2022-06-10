package decimal

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

// strconv.FormatFloat
// param1 浮点数
// param2 格式化模板、流程、类型（一般就是 'f'）
// param3 希望保留的小数位
// param4 32 或者 64（param1 是 float32 还是 float64 类型）

func TestStrconvFormatFloat(t *testing.T) {
	var f = 5.00

	// 需求：四舍五入 1 位小数
	s := math.Round(f*10) / 10

	t.Log(s)                                  // 得到 5
	t.Log(strconv.FormatFloat(s, 'f', 2, 64)) // 处理方式一
	t.Log(fmt.Sprintf("%.2f", s))             // 处理方式二
}
