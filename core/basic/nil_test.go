package basic

import (
	"fmt"
	"reflect"
	"testing"
)

/*
引出，初学者在学习类型零值的时候，可以了解到 指针类型、接口类型 的零值是 nil；但是没有进一步了解这之间的联系和区别就会遇到各式各样的问题

原理
- 接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V
	一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）
- 两个接口值比较时，会先比较 T，再比较 V；接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较
- 例，nil = interface{}(nil) != (*string)(nil)

直接说结论：比较并不难，一定要注意按照场景去区分，interface 的原理还是大头
- 具体 类型 和 nil 的比较：虽然关键字 nil 可以理解为 interface 类型 的 nil，但是只比较值，不看类型（仔细想想语言设计层面，简直就是废话）

- interface 类型 和 nil 的比较：T 是不是 nil，V 是不是 nil
- interface 类型 和 interface 类型 的比较：T 是不是相同，V 是不是相等
- interface 类型 和 具体类型 的比较：T 是不是相同，V 是不是相等
*/

func TestNil(t *testing.T) {
	interfaceNilTest()
	// specialNilTest()
	// specialTest()
}

func interfaceNilTest() {
	// 将 interface{} 替换成 error 也能得到相同的结果
	var nil1 interface{}
	var nil2 interface{} = nil
	var nil3 = (interface{})(nil)
	var nil4 = []interface{}{nil}[0]

	println(nil1 == nil)
	println(nil2 == nil)
	println(nil3 == nil)
	println(nil4 == nil)
}

// 很具有代表意义的测试样例
func specialNilTest() {
	var null1 interface{} = (*string)(nil)
	// var null1 = interface{}((*string)(nil)) // 和上面等效
	var null2 *string

	println(null1 == null2) // true
	println(null1 == nil)   // false
	println(null2 == nil)   // true

	println(null1 == (*string)(nil)) // true
	println(null2 == (*string)(nil)) // true


	// nilValue := reflect.ValueOf(null1)
	// println(nilValue.Type().String()) // *string
	// println(nilValue.Kind().String()) // ptr
}

func specialTest() {
	var i interface{}               // (nil nil) == nil → true
	var iPtr = &i                   // (*interface{} &I1) == nil → false（类型 和 值 都不匹配）
	var sPtr *string                // (*string nil) == nil → true（只看值，所以匹配）
	var sPtrWrap interface{} = sPtr // (*string sPtr) == nil → false（类型不匹配）

	fmt.Println(iPtr == (*interface{})(nil)) // false，类型相同，但是值不同（一个有，一个没有）

	fmt.Println(reflect.TypeOf(sPtrWrap)) // *string
	fmt.Println(reflect.TypeOf(iPtr))     // *interface {}
}
