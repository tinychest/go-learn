package basic

import (
	"reflect"
	"testing"
)

// 1、匹配到了 case，执行完就结束了（不接着走下边的 case）
// 2、没有匹配到 case，就走 default
// 3、case 后可以使用勾好隔开，定义多个值
func TestSwitch(t *testing.T) {
	switchTypeTest("") // string
	switchTypeTest(0)  // int
}

// 4、演示自动类型推断
func switchTypeTest(param interface{}) {
	switch t := param.(type) {
	case string:
		println(reflect.TypeOf(t).String())
	case int, int64:
		println(reflect.TypeOf(t).String())
	default:
		println("Unknown Type")
	}

	// 回忆一下，类型推断的类型转化
	// if afterTrans, ok := param.(string); ok {
	//     println(reflect.TypeOf(afterTrans).String())
	// }
}

// 替代多层的 if else
func TestReplaceIf(t *testing.T) {
	i := 0
	switch {
	case i > 0 && i <= 10:
		panic("123")
	case i > 10 && i <= 20:
		panic("456")
	}
}