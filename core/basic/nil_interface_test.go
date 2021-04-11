package basic

import (
	"reflect"
	"testing"
)

// 下边测试的结论：nil = interface{}(nil) != (*string)(nil)
// 有具体类型的指针 和 nil 作比较就会只比较值

/*
说的很好

接口(interface) 是对非接口值(例如指针，struct等)的封装，内部实现包含 2 个字段，类型 T 和 值 V
一个接口等于 nil，当且仅当 T 和 V 处于 unset 状态（T=nil，V is unset）

两个接口值比较时，会先比较 T，再比较 V
接口值与非接口值比较时，会先将非接口值尝试转换为接口值，再比较
*/

/*
展开说一下 go 中的 interface{}，一个为 interface{} 类型的 nil 值变量，如果你没有确定赋值的源头是写在脸上看到的 nil
那么将该变量与 nil 比较就要慎重了，因为无论你怎么传，赋值多少层，interface 底层始终会记录和保留最初的变量的类型
（当然，这里指定的是 nil 值，所以没有说存值，interface{} 就是会保留 类型 和 值 的）
*/
func TestInterface(t *testing.T) {
	// interface 类型的零值是 nil
	var null interface{}
	var null1 interface{} = nil
	var null2 = (interface{})(nil)
	var null3 = []interface{}{nil}[0]

	var null4 interface{} = (*string)(nil)
	var null5 = null4

	println(null == nil)
	println(null1 == nil)
	println(null2 == nil)
	println(null3 == nil)

	println(null4 == nil)
	println(null5 == nil)
	println(null5 == (*string)(nil))

	// 测试现象看完了，现在用反射来解释一下，所以可以很清楚的知道了，乍看他是一个 nil，实际它是 *string 的零值，即 nil 也是带有类型的
	nilValue := reflect.ValueOf(null4)
	println(nilValue.Type().String())
	println(nilValue.Kind().String())
	// 其他的 null 通过 reflect.ValueOf 方法返回的都是 nil
}
