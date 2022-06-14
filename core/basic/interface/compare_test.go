package _interface

import (
	"reflect"
	"testing"
)

// - 接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V
// - 两个接口值比较时，会先比较 T，再比较 V；接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较
// - 接口值的比较准则：T 是不是相同，V 是不是相等
//
// [nil compare]
// Go 中一个比较经典的问题就是 nil is not nil
//
// - 一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）
// - nil = interface{}(nil) != (*string)(nil) = nil
//
// [其他]
// 如果只希望得到一个 接口值 的 V 是否是 nil，可以通过 reflect.ValueOf(xxx).IsNil()

func TestCompare(t *testing.T) {
	// compareTest(t)
	compareNilTest(t)
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

func compareNilTest(t *testing.T) {
	var (
		null1 interface{} = (*string)(nil)
		// null1 = interface{}((*string)(nil))
		null2 *string
	)

	t.Log(null1 == nil)
	// 接口值之间的比较
	// 左边：interface（T = *string, V = nil）
	// 右边：interface（T = nil, V = nil）
	// 因此 false

	t.Log(null2 == nil)
	// 非接口值的比较（能进行比较说明类型相同）
	// 因此 true

	t.Log(null1 == null2)
	// 根据 接口值 和 非接口值 的比较原则
	// 左边：interface（T = nil, V = nil）
	// 等号左边是 接口值，右边不是，所以会将右边转成 interface{}，也就是
	// 右边：interface（T = *string, V = nil）
	// 因此 true

	// [验证]（reflect.ValueOf().String() 并不能得到期望的结果）
	t.Logf("T:%s, V:%v", reflect.TypeOf(null1).String(), null1)
	t.Logf("T:%s, V:%v", reflect.TypeOf(null2).String(), null2)

	// [拓展]
	t.Log(reflect.ValueOf(null1).IsNil())
	t.Log(reflect.ValueOf(null2).IsNil())
}
