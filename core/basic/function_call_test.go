package basic

import (
	"testing"
)

// 函数对应的结构体，由官方术语来说：函数 应该叫做 方法，结构体类型 应该叫做 方法的接收者（Receiver）

// 针对方法接收者的定义总结来说有如下几点（官方）
// 1、方法能够修改其接收者指向的值
// 2、可以避免在每次调用方法时复制该值（值传递）。若值的类型为大型结构体时，这样做会更加高效
// 3、通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用

// 语法点：如果方法的接收者是指针类型，我们应该说，结构体的指针类型实现了该接口，而不是该结构体类型（也可以说结构体类型的方法是结构体指针类型方法的一个子集）

type structForFun struct{}

func (s structForFun) nor() {}

func (s *structForFun) Ptr() {}

// TODO Invalid receiver type 'structForFunPtr' ('structForFunPtr' is a pointer type)
// type structForFunPtr *structForFun
// func (s structForFunPtr) abc() {}

func TestStructFunction(t *testing.T) {
	s := structForFun{}
	sPtr := &structForFun{}

	// 一、这都是都是基础部分，在 function_test.go 中都有例子实验，没什么好说的
	s.nor()
	s.Ptr()
	sPtr.nor()
	sPtr.Ptr()

	// 二、匿名结构体调用方法
	structForFun{}.nor()
	// structForFun{}.Ptr()  // 注意：编译不通过，匿名结构体不能进行指针类型的方法调用
	(&structForFun{}).nor()
	(&structForFun{}).Ptr()

	// 三、结构体方法的特殊调用形式
	structForFun.nor(s)
	// structForFun.nor(sPtr) // 注意：编译不通过，特殊的调用方法不支持结构体指针类型
	// structForFun.Ptr(s)    // 注意：goland 不提示编译错误，cannot use s (type structForFun) as type *structForFun in argument to (*structForFun).Ptr
	// structForFun.Ptr(sPtr) // 注意：goland 提示编译错误，  cannot use s (type structForFun) as type *structForFun in argument to (*structForFun).Ptr
	// (*structForFun).Ptr(s) // 注意：编译不通过，参数类型不对
	(*structForFun).Ptr(sPtr)

	// 四、不同的 Receiver 类型，对实际方法调用起的影响
	// 1、能不能真实影响真实值，不取决于实际调用方法的实例是否是指针类型，而是方法本身的定义
	// 2、为 person 类型绑定的方法，*person 的实例类型可以调用，为 *person 类型绑定的方法，person 的实例类型可以调用（go 底层做了转换）
	s.nor()
	s.Ptr()
	(&s).Ptr()
	sPtr.nor()
	(*sPtr).nor()
	sPtr.Ptr()

	// 五、*XxxStruct 可以访问 XxxStruct 的方法 和 字段
	// 接口类型就没有这种说法，*XxxInterface 不能直接调用 XxxInterface 的方法
}
