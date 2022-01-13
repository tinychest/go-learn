package number

import (
	"math"
	"testing"
)

// 在 Go 中直接定义常量表达式，xxx/0，编译时不报错，运行时 panic
// 但是如果是变量计算值，分母出现 0.0，并不会 panic（除以 0 依旧 panic），会得到一个特殊的 float64 类型的值
// 这个值通过 fmt 打印出来是 +Inf（为什么能打印出 +Inf，详见 go\src\strconv\ftoa.go/85）
func TestInf(t *testing.T) {
	// i1, i2 := 1, 0
	// i := i1 / i2 // 直接 panic
	// t.Log(i)

	f1, f2 := 1.0, -1.0
	nan1, nan2 := f1/0, f2/0
	t.Log(math.IsInf(nan1, 0))
	t.Log(math.IsInf(nan2, 0))
}
