package marshal

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

// 在 Go 中直接定义常量表达式，xxx/0，运行时直接 panic
// 但是如果是变量计算值，并不会 panic，会得到一个特殊的 float64 类型的值
// 这个值通过 fmt 打印出来是 +Inf

// 对于 json 来说如果是 "+Inf"，json 类库可以处理，但是直接返回 +Inf，是 类库无法处理的，将得到 panic
// 为什么 fmt 能打印出 +Inf，详见 go\src\strconv\ftoa.go/85
// json 类库如何返回 err，详见 go\src\encoding\json\encode.go/575
func TestMarshalInf(t *testing.T) {
	// i1, i2 := 1, 0
	// i := i1 / i2 // 直接 panic

	f1, f2 := 1.0, 2.0
	f := f1 / f2

	fmt.Println(f)
	fmt.Println(math.IsInf(f, 0))

	if result, err := json.Marshal(f); err != nil {
		panic(err)
	} else {
		fmt.Println(string(result))
	}
}
