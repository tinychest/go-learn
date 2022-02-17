package _reflect

import (
	. "reflect" // 介绍反射的包，必然进行包导入
	"testing"
)

// 我们希望有一个通用的方法，根据不同的参数类型动态决定返回值的类型
// 在 Java 中，有泛型，根据方法调用者类型，根据特定参数的类型，都能够在运行时，动态的决定类型
// 在 Go 中，想到类似的效果，语法上不支持，但是，开源框中总会遇到相同场景，参考
// 例，beego orm 的 QueryRows 就使用了一个很巧妙的做法，就是你把指定规则的容器传进来，我把数据放进去
func TestReflectApi(t *testing.T) {
	var i int
	var ip = &i
	var param interface{} = &ip

	// Value：数据类型在反射中的表现形式
	// Kind：类别
	// Type：具体类型

	// [Value ValueOf]
	// [Kind Value.Kind]
	// [Type Value.Type]
	value := ValueOf(param)
	kind := value.Kind()
	typ := value.Type()
	t.Logf("[Kind：%s] [Type：%v] [Value：%v]\n", kind, typ, value)

	// [bool Value.IsNil]
	// [Value Value.Elem]：interface 真实的值, Ptr 解引用
	// [Value Value.Indirect]：核心就是调用 Elem 方法
	indirect := Indirect(value)
	t.Logf("[Kind：%s] [Type：%v] [Indirect：%v]\n", indirect.Kind(), indirect.Type(), indirect)

	// [Value New(Type)]：根据指定类型创建一个实例
	// [Value Append(Value, ..Value)]：如果确认了方法第一个参数是切片类型，就可以通过该方法向切片中添加元素
	// [Value Set(Value)]：修改 Value 底层的值
	// [Value Addr(Value)]：源码注释 - 通常用于获取指向 “struct 字段” 或 “slice 元素” 的指针，以便调用需要指针接收器的方法

	// [bool CanAddr()]：源码注释 - 《addressable 可寻址的，这是一个相当重要的概念》：切片的元素，可寻址数组的元素，可寻址结构体的字段 或 解引用指针的结果
	// [Swapper(slice interface{}) func(i int, j int)]：返回一个用于交换指定切片的元素的方法

	// 测试：int → false、*int → false、**int → false
	if value.CanAddr() {
		value.SetInt(2)
	} else {
		t.Log("NO NO NO")
	}

	// 测试：int → false、*int → true、**int → true（但是 SetValue 会报错：panic: reflect: call of reflect.Value.SetInt on ptr Value）
	if indirect.CanAddr() {
		indirect.SetInt(2)
	} else {
		t.Log("NO NO NO")
	}

	t.Logf("%v\n", *(param.(*int)))

	// [int Value.NumField]：Struct 中字段的数量
	// [Value Value.Field(int)]：Struct 中指定下标的字段

	// [interface{} Value.Interface()]：返回底层真实的值（如果实际是结构体中未导出的值，则会发生 panic）
	// [int64 Value.Int()]：底层的值是 整数 类型，则可以使用该方法获取真实的值
	// ...

	// [int Value.Len()]：Array, Chan, Map, Slice, String 的长度
	// [Value Value.Index(int)]：Array, Slice, String 中指定下标的元素

	// [Type Type.Elem]：Array, Chan, Map, Ptr, or Slice 中元素的类型
	// [bool Type.Implements(Type)]：是否实现了指定的接口类型
}
