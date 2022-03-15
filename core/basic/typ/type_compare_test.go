package typ

import (
	"reflect"
	"testing"
)

// - 接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V
//	一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）
// - 两个接口值比较时，会先比较 T，再比较 V；接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较
// - 在 interface{} 和 各种 * 类型 与 预声明标识符 nil 的比较，不看类型，只看值是否是 nil（好像有点废话）
// - 涉及和 interface{} 类型包裹 的比较，既看类型，也看值：T 是不是相同，V 是不是相等

func TestTypeCompare(t *testing.T) {
	compareNilTest(t)
	compareTest(t)
}

// 很具有代表意义的测试样例
// nil = interface{}(nil) != (*string)(nil) = nil
func compareNilTest(t *testing.T) {
	var null1 interface{} = (*string)(nil)
	// var null1 = interface{}((*string)(nil)) // 和上面等效
	var null2 *string

	t.Log(null1 == null2) // true
	t.Log(null2 == nil)   // true
	t.Log(null1 == nil)   // false

	t.Log(null1 == (*string)(nil)) // true
	t.Log(null2 == (*string)(nil)) // true

	// nilValue := reflect.ValueOf(null1)
	// t.Log(nilValue.Type().String()) // *string
	// t.Log(nilValue.Kind().String()) // ptr
}

func compareTest(t *testing.T) {
	var i interface{}               // <nil nil> == nil → true
	var iPtr = &i                   // <*interface{} &I1> == nil → false（值不等于 nil）
	var sPtr *string                // <*string nil> == nil → true（值等于 nil）
	var sPtrWrap interface{} = sPtr // <*string nil> == nil → false（被 interface{} 包裹，就要看类型）

	// false，已经和 interface{} 没有关闭了，这里是 *interface{} 类型，这里是类型相同（可以比较），但是值不同（一个有，一个没有）
	// 不好理解的话，就将 *interface{} 改为 string
	t.Log(iPtr == (*interface{})(nil))

	t.Log(reflect.TypeOf(sPtrWrap)) // *string
	t.Log(reflect.TypeOf(iPtr))     // *interface {}
}