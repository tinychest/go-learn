package basic

import (
	"reflect"
	"testing"
)

// 1、匹配到了 case，执行完就结束了（不接着走下边的 case）
// 2、没有匹配到 case，就走 default
func TestSwitch(t *testing.T) {
	switchTypeTest("") // string
	switchTypeTest(0)  // int
}

// 3、演示自动类型推断
func switchTypeTest(param interface{}) {
	switch afterTrans := param.(type) {
	case string:
		println(reflect.TypeOf(afterTrans).String())
	case int:
		println(reflect.TypeOf(afterTrans).String())
	}

	// 回忆一下，类型推断的类型转化
	// if afterTrans, ok := param.(string); ok {
	//     println(reflect.TypeOf(afterTrans).String())
	// }
}
