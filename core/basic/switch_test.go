package basic

import (
	"testing"
)

/*
- 匹配到了 case，执行完 case，select 代码块就结束了（和别的高级语言不同）
- case 后可以使用逗号隔开，定义多个值（和别的高级语言不同）
- 没有匹配到 case，就走 default
- 可以通过在 case 语句体的最后添加 fallthrough 关键字，来起来继续向下走的作用（无视下面 case 的条件，如果还有 fallthrough 就再接着走，default 也同样对待）
*/

func TestSwitch(t *testing.T) {
	switchTypeTest(t, nil)              // nil match
	switchTypeTest(t, interface{}(nil)) // nil match
	switchTypeTest(t, (*string)(nil))   // nil match
	switchTypeTest(t, 1)                // interface{} match
	switchTypeTest(t, "")               // interface{} match
	switchTypeTest(t, []int{})          // interface{} match

}

// 自动类型推断（这种语法只支持在 switch 中使用）
// - nil：是一个预声明的标识符，没有默认类型，interface{} 都无法匹配，需要特殊处理（更多有关 nil 的内容，详见 nil_test）
// - interface{}：除了 nil 的万能匹配类型
// - types 包下的类型是不符合这里语境的（都是特殊结构体类型）
func switchTypeTest(t *testing.T, p interface{}) {
	switch p.(type) {
	case nil:
		t.Log("nil match")
	case interface{}:
		t.Log("interface{} match")
	default:
		t.Log("default match")
	}
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
