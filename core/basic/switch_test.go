package basic

import (
	"reflect"
	"testing"
)

/*
【语法】
- 匹配到了 case，执行完对应就结束了；和别的高级语言不同，是不会接着走下边的 case 的
- case 后可以使用逗号隔开，定义多个值
- 没有匹配到的 case，就走 default

- 可以通过在 case 语句体的最后添加 fallthrough 关键字，来起来继续向下走的作用（无视下面 case 的条件，如果还有 fallthrough 就再接着走，default 也同样对待）
*/

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
	n := 0
	switch {
	case n > 0 && n <= 10:
		println(123)
	case n > 10 && n <= 20:
		println(456)
	}
}