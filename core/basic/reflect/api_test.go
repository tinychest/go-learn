package _reflect

import (
	"fmt"
	. "reflect" // 介绍反射的包，必然进行包导入
	"testing"
)

// 我们希望有一个通用的方法，根据不同的参数返回不同类型的数据
// 在 Java 中，有泛型，根据方法调用者类型，根据特定参数的类型，都能够在运行时，动态的决定类型
// 在 Go 中，虽然不是很灵活，但是能做到根据一个表示要返回什么类型的参数来决定具体返回的类型，但是只要涉及到”通用“二字那么在声明上类型肯定是 interface{} 了
// 所以 beego orm 的 QueryRows 就使用了一个很巧妙的做法，就是你传容器进来，我把数据放进去

// 经验之谈1：反射中的递归要注意一个问题，因为反射就会得到 Value，这就是一个结构体
// 经验之谈2：获取反射对象（Value）的真实值，只要调用 Value.类型() 方法就可以，这里边有一个比较重要的方法 Interface()
func TestReflectApi(t *testing.T) {
	var i int
	var iPtr = &i
	var iPtrPtr = &iPtr

	tt(iPtrPtr)

	notUseButToCompile(i, iPtr, iPtrPtr)
}

func notUseButToCompile(...interface{}) {}

func tt(param interface{}) {
	// Value：数据类型在反射中的表现形式
	// Kind：类别
	// Type：具体类型

	// [Value ValueOf]
	// [Kind Value.Kind]
	// [Type Value.Type]
	value := ValueOf(param)
	kind := value.Kind()
	typ := value.Type()
	fmt.Printf("[Kind：%s] [Type：%v] [Value：%v]\n", kind, typ, value)

	// [bool Value.IsNil]
	// [Value Value.Elem]：interface 真实的值, Ptr 解引用
	// [Value Value.Indirect]：核心就是调用 Elem 方法
	indirect := Indirect(value)
	fmt.Printf("[Kind：%s] [Type：%v] [Indirect：%v]\n", indirect.Kind(), indirect.Type(), indirect)

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
		println("NO NO NO")
	}

	// 测试：int → false、*int → true、**int → true（但是 SetValue 会报错：panic: reflect: call of reflect.Value.SetInt on ptr Value）
	if indirect.CanAddr() {
		indirect.SetInt(2)
	} else {
		println("NO NO NO")
	}

	fmt.Printf("%v\n", *(param.(*int)))

	// [int Value.NumField]：Struct 中字段的数量
	// [Value Value.Field(int)]：Struct 中指定下标的字段
	// [interface{} Value.Interface()]：把底层的值作为 interface 类型返回（底层值如果为 nil 或者 结构体中未导出的值，则会发生空指针异常）

	// [int Value.Len()]：Array, Chan, Map, Slice, String 的长度
	// [Value Value.Index(int)]：Array, Slice, String 中指定下标的元素

	// [Type Type.Elem]：Array, Chan, Map, Ptr, or Slice 中元素的类型
	// [bool Type.Implements(Type)]：是否实现了指定的接口类型
}

func UnsafePackApi() {
	// [unsafe.Sizeof(interface{})]：计算出一个数据类型实例需要占用的字节数
}
